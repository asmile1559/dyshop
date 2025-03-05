package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/asmile1559/dyshop/app/checkout/biz/dal"
	"github.com/asmile1559/dyshop/app/checkout/biz/model"
	pbcheckout "github.com/asmile1559/dyshop/pb/backend/checkout"
)

// CheckoutService 结算服务
type CheckoutService struct {
	ctx context.Context
}

// NewCheckoutService 创建结算服务实例
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Checkout 处理结算请求
func (s *CheckoutService) Run(req *pbcheckout.CheckoutReq) (*pbcheckout.CheckoutResp, error) {
	transactionId := generateTransactionID()
	address := req.Address

	// 计算订单总商品数
	totalQuantity := 0
	for _, product := range req.Products {
		totalQuantity += int(product.Quantity)
	}

	// 创建订单记录
	orderRecord := model.OrderRecord{
		OrderID:       req.OrderId,
		UserID:        req.UserId,
		TransactionID: transactionId, 
		Recipient:     address.Recipient, // 这里如果有用户信息，可以填充
		Phone:         address.Phone, // 这里如果有电话号码，可以填充
		Province:      address.Province,
		City:          address.City,
		District:      address.District,
		Street:        address.Street,
		FullAddress:   address.FullAddress,
		TotalQuantity: totalQuantity,
		TotalPrice:    req.OrderPrice,
		Postage:       req.OrderPostage,
		FinalPrice:    req.OrderFinalPrice,
		CreatedAt:     time.Now(),
	}

	// 存储订单记录
	if err := dal.DB.Create(&orderRecord).Error; err != nil {
		return nil, fmt.Errorf("failed to save order: %v", err)
	}

	// 存储订单商品
	for _, product := range req.Products {
		orderItem := model.OrderItem{
			OrderID:       req.OrderId,
			TransactionID: transactionId, 
			ProductID:     product.ProductId,
			ProductImg:    product.ProductImg,
			ProductName:   product.ProductName,
			SpecName:      product.ProductSpec.Name,
			SpecPrice:     parsePrice(product.ProductSpec.Price),
			Quantity:      int(product.Quantity),
			Postage:       product.Postage,
			Currency:      product.Currency,
		}

		if err := dal.DB.Create(&orderItem).Error; err != nil {
			return nil, fmt.Errorf("failed to save order item: %v", err)
		}
	}

	// 返回结算响应
	return &pbcheckout.CheckoutResp{
		OrderId:       req.OrderId,
		TransactionId: transactionId,
	}, nil
}

func generateTransactionID() string {
	now := time.Now().UnixNano()
	rand.Seed(now)
	randomPart := rand.Intn(1000000)
	return fmt.Sprintf("TXN%d%06d", now, randomPart)
}

// 解析价格（从字符串转换为浮点数）
func parsePrice(priceStr string) float64 {
	var price float64
	_, err := fmt.Sscanf(priceStr, "%f", &price)
	if err != nil {
		return 0.0 // 默认返回0，避免错误
	}
	return price
}
