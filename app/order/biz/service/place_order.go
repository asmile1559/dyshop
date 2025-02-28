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
	s := &PlaceOrderService{ctx: c, DB: db.DB}
	go s.startDeleteUnpaidOrdersTask()
	return s
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
	modelAddress := model.PrePaidAddress{
		StreetAddress: address.GetStreetAddress(),
		City:          address.GetCity(),
		State:         address.GetState(),
		Country:       address.GetCountry(),
		ZipCode:       address.GetZipCode(),
	}

	orderItems := make([]model.PrePaidOrderItem, len(req.GetOrderItems()))
	for i, item := range req.GetOrderItems() {
		cartItem := item.GetItem()
		if cartItem == nil {
			return nil, fmt.Errorf("cart item is required")
		}
		orderItems[i] = model.PrePaidOrderItem{
			ProductId: uint64(cartItem.GetProductId()),
			Quantity:  int(cartItem.GetQuantity()),
			Cost:      float64(item.GetCost()),
		}
	}

	newOrder := model.PrePaidOrder{
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
		// 创建预支付订单
		if err := tx.Create(&newOrder).Error; err != nil {
			logrus.Error("Failed to create prepaid order:", err)
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

func (s *PlaceOrderService) startDeleteUnpaidOrdersTask() {
	// 创建一个独立于外部 ctx 的新根上下文，并设置超时
	deleteCtx, deleteCancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer deleteCancel()

	interval := 2 * time.Minute // 每2分钟检查一次
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.deleteUnpaidOrders(deleteCtx, interval)
			return
		case <-deleteCtx.Done():
			logrus.Info("Stopping deleteUnpaidOrders task")
			return
		}
	}
}

func (s *PlaceOrderService) deleteUnpaidOrders(ctx context.Context, interval time.Duration) {
	cutoffTime := time.Now().Add(-1 * interval) // 超过2分钟未支付的订单
	var unpaidOrders []model.PrePaidOrder
	if err := s.DB.WithContext(ctx).Where("paid = ? AND created_at < ?", false, cutoffTime).Find(&unpaidOrders).Error; err != nil {
		logrus.Error("Failed to find unpaid orders:", err)
		return
	}
	for _, order := range unpaidOrders {
		logrus.Infof("Found unpaid order with ID: %v", order.ID)
	}

	/*for _, order := range unpaidOrders {//删除未支付且超时的订单
		if err := s.DB.WithContext(ctx).Delete(&order).Error; err != nil {
			logrus.Error("Failed to delete unpaid order:", err)
		} else {
			logrus.Infof("Delete unpaid order with ID: %v", order.ID)
		}
	}*/
}
