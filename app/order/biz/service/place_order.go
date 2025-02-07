package service

import (
	"context"
	"github.com/asmile1559/dyshop/app/order/biz/model"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type PlaceOrderService struct {
	ctx context.Context
	DB  *gorm.DB
}

func NewPlaceOrderService(c context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: c, DB: db.DB}
}

/*
func (s *PlaceOrderService) Run(req *pborder.PlaceOrderReq) (*pborder.PlaceOrderResp, error) {
	// TODO: finish your business code...
	//

	return &pborder.PlaceOrderResp{
		Order: &pborder.OrderResult{OrderId: "1"},
	}, nil

}*/

/*
func (s *PlaceOrderService) Run(req *pborder.PlaceOrderReq) (*pborder.PlaceOrderResp, error) {
	// 假设这里有一些业务逻辑处理代码...
	// 例如：计算总费用、验证库存、更新数据库等...

	// 示例数据填充，实际使用时应该替换为从请求或其他服务获取的数据
	address := &pborder.Address{
		StreetAddress: "123 Main St",
		City:          "Anytown",
		State:         "CA",
		Country:       "USA",
		ZipCode:       "90210",
	}
	orderItems := []*pborder.OrderItem{
		{
			Item: &pbcart.CartItem{ProductId: 101, Quantity: 2},
			Cost: 49.98,
		},
		// 可以添加更多的OrderItem实例...
	}

	resp := &pborder.PlaceOrderResp{
		UserId:       123,   // 假设用户ID是123
		UserCurrency: "USD", // 用户货币假设为美元
		Address:      address,
		Email:        "user@example.com",
		OrderItems:   orderItems,
	}

	// 返回响应
	return resp, nil
}*/

func generateUniqueOrderID() uint64 {
	uniqueID := uint64(time.Now().UnixNano())
	return uniqueID
}

func (s *PlaceOrderService) Run(req *pborder.PlaceOrderReq) (*pborder.PlaceOrderResp, error) {
	orderID := generateUniqueOrderID()

	address := model.Address{
		StreetAddress: req.GetStreetAddress(),
		City:          req.GetCity(),
		State:         req.GetState(),
		Country:       req.GetCountry(),
		ZipCode:       req.GetZipCode(),
	}

	orderItems := make([]model.OrderItem, len(req.GetOrderItems()))
	for i, item := range req.GetOrderItems() {
		orderItems[i] = model.OrderItem{
			ProductID: uint64(item.GetItem.GetProductId()),
			Quantity:  int(item.GetItem.GetQuantity()),
			Cost:      float64(item.GetCost()),
		}
	}

	newOrder := model.Order{
		ID:           orderID,
		UserId:       req.GetUserId(),
		UserCurrency: req.GetUserCurrency(),
		Address:      address,
		Email:        req.GetEmail(),
		OrderItems:   orderItems,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		//创建订单
		if err := tx.Create(&newOrder).Error; err != nil {
			logrus.Error("Failed to create order:", err)
			return err
		}
		//更新地址表中的OrderID
		address.OrderID = newOrder.ID
		if err := tx.Save(&address).Error; err != nil {
			logrus.Error("Failed to save address:", err)
			return err
		}
		//创建订单项
		for i := range orderItems {
			orderItems[i].OrderID = newOrder.ID
			if err := tx.Create(&orderItems[i]).Error; err != nil {
				logrus.Error("Failed to create orderItem:", err)
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	resp := &pborder.PlaceOrderResp{
		Order: &pborder.OrderResult{OrderId: orderID},
	}
	return resp, nil
}
