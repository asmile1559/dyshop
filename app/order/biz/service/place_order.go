package service

import (
	"context"
	"fmt"
	"github.com/asmile1559/dyshop/app/order/biz/model"
	"github.com/asmile1559/dyshop/app/order/utils/db"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type PlaceOrderService struct {
	ctx context.Context
	DB  *gorm.DB
}

func NewPlaceOrderService(c context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: c, DB: db.DB}
}

func generateUniqueOrderID() uint64 {
	uniqueID := uint64(time.Now().UnixNano())
	return uniqueID
}

func (s *PlaceOrderService) Run(req *pborder.PlaceOrderReq) (*pborder.PlaceOrderResp, error) {
	orderID := generateUniqueOrderID()

	address := req.GetAddress()
	if address == nil {
		return nil, fmt.Errorf("address is nil")
	}
	modelAddress := model.Address{
		StreetAddress: address.GetStreetAddress(),
		City:          address.GetCity(),
		State:         address.GetState(),
		Country:       address.GetCountry(),
		ZipCode:       address.GetZipCode(),
	}

	orderItems := make([]model.OrderItem, len(req.GetOrderItems()))
	for i, item := range req.GetOrderItems() {
		cartItem := item.GetItem()
		if cartItem == nil {
			return nil, fmt.Errorf("cart item is required")
		}
		orderItems[i] = model.OrderItem{
			ProductId: uint64(cartItem.GetProductId()),
			Quantity:  int(cartItem.GetQuantity()),
			Cost:      float64(item.GetCost()),
		}
	}

	newOrder := model.Order{
		ID:           orderID,
		UserId:       uint64(req.GetUserId()),
		UserCurrency: req.GetUserCurrency(),
		Address:      modelAddress,
		Email:        req.GetEmail(),
		Paid:         false,
		OrderItems:   orderItems,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		PaidAt:       time.Now(),
	}

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		//创建订单
		if err := tx.Create(&newOrder).Error; err != nil {
			logrus.Error("Failed to create order:", err)
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	resp := &pborder.PlaceOrderResp{
		Order: &pborder.OrderResult{
			OrderId: strconv.FormatUint(orderID, 10), // 使用 strconv 包将 uint64 转换为字符串
		},
	}
	return resp, nil
}
