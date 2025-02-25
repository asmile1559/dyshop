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
	"google.golang.org/grpc/credentials/insecure"
)

// initServiceViper 读取指定配置文件
func initServiceViper(configPath string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}

// CheckoutService 结算服务（不依赖 etcd，只使用纯 gRPC）
type CheckoutService struct {
	ctx        context.Context
	paymentCli pbpayment.PaymentServiceClient
	orderCli   pborder.OrderServiceClient
}

// NewCheckoutService 创建结算服务实例，分别从 payment 和 order 的配置文件中读取服务地址
func NewCheckoutService(ctx context.Context) *CheckoutService {
	// 读取 payment 服务的配置文件
	paymentViper, err := initServiceViper("/root/dyshop/app/payment/conf/config.yaml")
	if err != nil {
		logrus.Fatalf("读取 payment 配置文件失败: %v", err)
	}
	paymentPort := paymentViper.GetString("server.port")
	// 假设 payment 服务部署在本机
	paymentAddr := "localhost:" + paymentPort

	// 读取 order 服务的配置文件
	orderViper, err := initServiceViper("/root/dyshop/app/order/conf/config.yaml")
	if err != nil {
		logrus.Fatalf("读取 order 配置文件失败: %v", err)
	}
	orderPort := orderViper.GetString("server.port")
	// 假设 order 服务部署在本机
	orderAddr := "localhost:" + orderPort

	// 建立与 payment 服务的 gRPC 连接
	paymentConn, err := grpc.Dial(paymentAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("连接 payment 服务失败: %v", err)
	}

	// 建立与 order 服务的 gRPC 连接
	orderConn, err := grpc.Dial(orderAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("连接 order 服务失败: %v", err)
	}

	return &CheckoutService{
		ctx:        ctx,
		paymentCli: pbpayment.NewPaymentServiceClient(paymentConn),
		orderCli:   pborder.NewOrderServiceClient(orderConn),
	}
}

// Run 处理结算请求
func (s *CheckoutService) Run(req *pbcheckout.CheckoutReq) (*pbcheckout.CheckoutResp, error) {
	// 校验请求
	if err := validateCheckoutRequest(req); err != nil {
		return nil, err
	}

	// 从 order 服务获取订单信息
	orderResp, err := s.getOrderFromOrderService(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("获取订单失败: %v", err)
	}

	// 如果订单列表为空
	if len(orderResp.Orders) == 0 {
		return nil, errors.New("未找到订单")
	}
	order := orderResp.Orders[0] // 取第一笔订单
	orderID := order.OrderId

	fmt.Println("订单测试：", order.OrderItems)

	// 计算订单金额
	totalAmount, err := calculateTotalAmount(order.OrderItems)
	if err != nil {
		return nil, err
	}
	fmt.Println("订单金额：", totalAmount)

	// 调用 payment 服务进行支付
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
		return nil, errors.New("支付失败")
	}

	// 存储订单信息到数据库
	record := model.OrderRecord{
		OrderID:       orderID,
		UserID:        req.UserId,
		Amount:        totalAmount,
		TransactionID: paymentResp.TransactionId,
		Status:        "SUCCESS",
		CreatedAt:     time.Now(),
	}
	if err := dal.DB.Create(&record).Error; err != nil {
		return nil, fmt.Errorf("保存订单失败: %v", err)
	}

	// 返回订单信息
	return &pbcheckout.CheckoutResp{
		OrderId:       orderID,
		TransactionId: paymentResp.TransactionId,
	}, nil
}

// getOrderFromOrderService 获取订单信息
func (s *CheckoutService) getOrderFromOrderService(userId uint32) (*pborder.ListOrderResp, error) {
	resp, err := s.orderCli.ListOrder(context.TODO(), &pborder.ListOrderReq{UserId: userId})
	if err != nil {
		return nil, fmt.Errorf("获取订单数据失败: %v", err)
	}
	fmt.Printf("订单列表: %+v\n", resp)
	return resp, nil
}

// validateCheckoutRequest 校验结算请求
func validateCheckoutRequest(req *pbcheckout.CheckoutReq) error {
	if req.UserId == 0 || req.Firstname == "" || req.Lastname == "" || req.Email == "" {
		return errors.New("缺少必要字段")
	}
	if req.Address.StreetAddress == "" || req.Address.City == "" || req.Address.ZipCode == "" {
		return errors.New("地址无效")
	}
	if req.CreditCard.CreditCardNumber == "" || req.CreditCard.CreditCardCvv == 0 {
		return errors.New("信用卡信息无效")
	}
	return nil
}

// calculateTotalAmount 计算订单总金额
func calculateTotalAmount(items []*pborder.OrderItem) (float64, error) {
	var total float64
	for _, item := range items {
		total += float64(item.Cost)
	}
	return total, nil
}
