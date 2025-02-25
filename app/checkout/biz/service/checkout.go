package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/asmile1559/dyshop/app/checkout/biz/dal"
	"github.com/asmile1559/dyshop/app/checkout/biz/model"
	pbcheckout "github.com/asmile1559/dyshop/pb/backend/checkout"
	pbpayment "github.com/asmile1559/dyshop/pb/backend/payment"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials/insecure"
	"github.com/asmile1559/dyshop/utils/registryx"
)

// CheckoutService 结算服务，通过 etcd 发现依赖服务
type CheckoutService struct {
	ctx        context.Context
	// 通过 etcd 服务发现获得 Payment 和 Order 客户端
	paymentCli pbpayment.PaymentServiceClient
	orderCli   pborder.OrderServiceClient
}

// NewCheckoutService 创建结算服务实例，使用 registryx.DiscoverEtcdServices 发现 Payment 和 Order 服务
func NewCheckoutService(ctx context.Context) *CheckoutService {
	etcdEndpoints := viper.GetStringSlice("etcd.endpoints")

	// 发现 Payment 服务
	paymentClient, _, err := registryx.DiscoverEtcdServices[pbpayment.PaymentServiceClient](
		etcdEndpoints,
		"/services/payment",
		func(conn grpc.ClientConnInterface) pbpayment.PaymentServiceClient {
			return pbpayment.NewPaymentServiceClient(conn)
		},
	)
	if err != nil {
		logrus.Fatalf("Failed to discover payment service: %v", err)
	}
	// 如果需要，请在服务关闭时关闭 paymentConn

	// 发现 Order 服务
	orderClient, _, err := registryx.DiscoverEtcdServices[pborder.OrderServiceClient](
		etcdEndpoints,
		"/services/order",
		func(conn grpc.ClientConnInterface) pborder.OrderServiceClient {
			return pborder.NewOrderServiceClient(conn)
		},
	)
	if err != nil {
		logrus.Fatalf("Failed to discover order service: %v", err)
	}
	// 同样，orderConn 关闭由调用方负责

	return &CheckoutService{
		ctx:        ctx,
		paymentCli: paymentClient,
		orderCli:   orderClient,
	}
}

// Run 处理结算请求
func (s *CheckoutService) Run(req *pbcheckout.CheckoutReq) (*pbcheckout.CheckoutResp, error) {
	// 校验请求
	if err := validateCheckoutRequest(req); err != nil {
		return nil, err
	}

	// 通过 Order 服务获取订单信息
	orderResp, err := s.getOrderFromOrderService(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders: %v", err)
	}
	if len(orderResp.Orders) == 0 {
		return nil, errors.New("no orders found")
	}
	order := orderResp.Orders[0] // 取第一笔订单
	orderID := order.OrderId

	// 计算订单总金额
	totalAmount, err := calculateTotalAmount(order.OrderItems)
	if err != nil {
		return nil, err
	}
	fmt.Println("Order total:", totalAmount)

	// 调用 Payment 服务进行支付
	paymentResp, err := s.paymentCli.Charge(context.TODO(), &pbpayment.ChargeReq{
		Amount: float32(totalAmount),
		CreditCard: &pbpayment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
		},
		OrderId: orderID,
		UserId:  req.UserId,
	})
	if err != nil || paymentResp.TransactionId == "" {
		return nil, errors.New("payment failed")
	}

	// 存储订单记录到数据库
	record := model.OrderRecord{
		OrderID:       orderID,
		UserID:        req.UserId,
		Amount:        totalAmount,
		TransactionID: paymentResp.TransactionId,
		Status:        "SUCCESS",
		CreatedAt:     time.Now(),
	}
	if err := dal.DB.Create(&record).Error; err != nil {
		return nil, fmt.Errorf("failed to save order: %v", err)
	}

	return &pbcheckout.CheckoutResp{
		OrderId:       orderID,
		TransactionId: paymentResp.TransactionId,
	}, nil
}

// getOrderFromOrderService 调用 Order 服务的 ListOrder 方法获取订单数据
func (s *CheckoutService) getOrderFromOrderService(userId uint32) (*pborder.ListOrderResp, error) {
	resp, err := s.orderCli.ListOrder(context.TODO(), &pborder.ListOrderReq{UserId: userId})
	if err != nil {
		return nil, fmt.Errorf("failed to get order data: %v", err)
	}
	fmt.Printf("Orders: %+v\n", resp)
	return resp, nil
}

func validateCheckoutRequest(req *pbcheckout.CheckoutReq) error {
	if req.UserId == 0 || req.Firstname == "" || req.Lastname == "" || req.Email == "" {
		return errors.New("missing required fields")
	}
	if req.Address.StreetAddress == "" || req.Address.City == "" || req.Address.ZipCode == "" {
		return errors.New("invalid address")
	}
	if req.CreditCard.CreditCardNumber == "" || req.CreditCard.CreditCardCvv == 0 {
		return errors.New("invalid credit card info")
	}
	return nil
}

func calculateTotalAmount(items []*pborder.OrderItem) (float64, error) {
	var total float64
	for _, item := range items {
		total += float64(item.Cost)
	}
	return total, nil
}
