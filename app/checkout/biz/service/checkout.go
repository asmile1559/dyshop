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
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// CheckoutService 结算服务
type CheckoutService struct {
	ctx        context.Context
	paymentCli pbpayment.PaymentServiceClient
	orderCli   pborder.OrderServiceClient
}

// NewCheckoutService 创建结算服务实例
func NewCheckoutService(c context.Context) *CheckoutService {
	// 连接支付服务
	paymentViper, err := InitViper("/home/djj/devel/dyshop/app/payment/conf/config.yaml")
    if err != nil {
        logrus.Fatalf("Error reading order service config file, %s", err)
    }

    paymentPort := paymentViper.GetString("server.port")
	conn, err := grpc.Dial("localhost:"+paymentPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to payment service: " + err.Error())
	}

	return &CheckoutService{
		ctx:        c,
		paymentCli: pbpayment.NewPaymentServiceClient(conn),
	}
}

// Run 处理结算请求
func (s *CheckoutService) Run(req *pbcheckout.CheckoutReq) (*pbcheckout.CheckoutResp, error) {
	// 校验请求
	if err := validateCheckoutRequest(req); err != nil {
		return nil, err
	}

	// 获取订单
	orderResp, err := getOrderFromOrderService()
	if err != nil {
		return nil, fmt.Errorf("获取订单失败: %v", err)
	}

	// 取出订单
	if len(orderResp.Orders) == 0 {
		return nil, errors.New("未找到订单")
	}
	order := orderResp.Orders[0] // 取第一笔订单
	orderID := order.OrderId

	// 计算订单金额
	totalAmount, err := calculateTotalAmount(order.OrderItems)
	fmt.Println("订单金额：", totalAmount)
	if err != nil {
		return nil, err
	}

	// 调用支付服务
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
		total += float64(item.Cost) // Cost 大写，float32 转 float64
	}

	return total, nil
}

// getOrderFromOrderService 获取订单信息
func getOrderFromOrderService() (*pborder.ListOrderResp, error) {
	orderViper, err := InitViper("/home/djj/devel/dyshop/app/order/conf/config.yaml")
    if err != nil {
        logrus.Fatalf("Error reading order service config file, %s", err)
    }

    orderPort := orderViper.GetString("server.port")

	cc, err := grpc.Dial("localhost:"+orderPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        logrus.Fatalf("连接 OrderService 失败: %v", err)
    }
    defer cc.Close() // 确保在 main 结束时关闭连接

    // 创建 OrderService 客户端
    cli := pborder.NewOrderServiceClient(cc)

    // 发送请求获取订单列表
    resp, err := cli.ListOrder(context.TODO(), &pborder.ListOrderReq{UserId: 1})
    if err != nil {
		return nil, fmt.Errorf("获取订单数据失败: %v", err)
	}	

    fmt.Printf("订单列表: %+v\n", resp)
	return resp, nil
}

func InitViper(configPath string) (*viper.Viper, error) {
    v := viper.New()
    v.SetConfigFile(configPath)
    if err := v.ReadInConfig(); err != nil {
        return nil, err
    }
    return v, nil
}

