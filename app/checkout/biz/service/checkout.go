package service

import (
	"context"
	"fmt"
	"time"
	"math/rand"
	"github.com/asmile1559/dyshop/app/checkout/biz/dal"
	"github.com/asmile1559/dyshop/app/checkout/biz/model"
	pbcheckout "github.com/asmile1559/dyshop/pb/backend/checkout"
)

// CheckoutService 结算服务
type CheckoutService struct {
	ctx        context.Context
}

// NewCheckoutService 创建结算服务实例
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Checkout 处理结算请求
func (s *CheckoutService) Run(req *pbcheckout.CheckoutReq) (*pbcheckout.CheckoutResp, error) {

	transactionId := generateTransactionID()
	address := req.Address

	// 存订单记录
    orderRecord := model.OrderRecord{
        OrderID:       req.OrderId,
        UserID:        req.UserId,
        TransactionID: transactionId, // 记录 TransactionID
        Recipient:     req.Lastname + req.Firstname,
        Phone:         "12345678901",
        Province:      address.State,
        City:          address.City,
        District:      "海淀区",
        Street:        address.StreetAddress,
        FullAddress:   "北京北京市海淀区知春路甲48号抖音视界",
        TotalQuantity: 3,
        TotalPrice:    58.59,
        Postage:       10,
        FinalPrice:    68.59,
        CreatedAt:     time.Now(),
    }

    if err := dal.DB.Create(&orderRecord).Error; err != nil {
        return nil, fmt.Errorf("failed to save order: %v", err)
    }

	orderItem := model.OrderItem{
		OrderID:       req.OrderId,
		TransactionID: transactionId, // 记录 TransactionID
		ProductID:     "1",
		ProductImg:    "/static/src/product/bearcookie.webp",
		ProductName:   "超级无敌好吃的小熊饼干",
		SpecName:      "500g装",
		SpecPrice:     18.80,
		Quantity:      2,
		Postage:       10,
		Currency:      "CNY",
	}

	if err := dal.DB.Create(&orderItem).Error; err != nil {
		return nil, fmt.Errorf("failed to save order item: %v", err)
	}

	orderItem = model.OrderItem{
		OrderID:       req.OrderId,
		TransactionID: transactionId, // 记录 TransactionID
		ProductID:     "2",
		ProductImg:    "/static/src/product/bearsweet.webp",
		ProductName:   "超级无敌好吃的小熊软糖值得品尝大力推荐",
		SpecName:      "9分软",
		SpecPrice:     20.99,
		Quantity:      1,
		Postage:       0,
		Currency:      "",
	}

	if err := dal.DB.Create(&orderItem).Error; err != nil {
		return nil, fmt.Errorf("failed to save order item: %v", err)
	}


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
