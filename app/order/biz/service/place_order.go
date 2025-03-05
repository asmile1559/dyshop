package service

import (
	"context"
	"strings"
	"time"

	"fmt"

	"github.com/asmile1559/dyshop/app/order/biz/model"
	"github.com/asmile1559/dyshop/app/order/utils/db"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

func generateUniqueOrderID() uint32 {
	uniqueID := uint32(time.Now().UnixNano())
	return uniqueID
}

func (s *PlaceOrderService) Run(req *pborder.PlaceOrderReq) (*pborder.PlaceOrderResp, error) {
	orderID := generateUniqueOrderID()

	pids := []string{}
	for _, pid := range req.GetProductIds() {
		pids = append(pids, fmt.Sprint(pid))
	}
	newOrder := model.Order{
		Model:      gorm.Model{ID: uint(orderID)},
		UserId:     req.GetUserId(),
		AddressId:  req.GetAddressId(),
		Price:      req.GetPrice(),
		ProductIDs: strings.Join(pids, ","),
	}

	err := db.DB.Create(&newOrder).Error
	if err != nil {
		return nil, err
	}
	resp := &pborder.PlaceOrderResp{
		OrderId: orderID,
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
		case <-deleteCtx.Done():
			logrus.Info("Stopping deleteUnpaidOrders task")
			return
		}
	}
}

func (s *PlaceOrderService) deleteUnpaidOrders(ctx context.Context, interval time.Duration) {
	cutoffTime := time.Now().Add(-1 * interval) // 超过2分钟未支付的订单
	var unpaidOrders []model.Order
	if err := s.DB.WithContext(ctx).Where("paid = ? AND created_at < ?", false, cutoffTime).Find(&unpaidOrders).Error; err != nil {
		logrus.Error("Failed to find unpaid orders:", err)
		return
	}
	for _, order := range unpaidOrders {
		logrus.Infof("Found unpaid order with ID: %v", order.ID)
	}

	for _, order := range unpaidOrders { // 删除未支付且超时的订单
		if err := s.DB.WithContext(ctx).Delete(&order).Error; err != nil {
			logrus.Error("Failed to delete unpaid order:", err)
		} else {
			logrus.Infof("Deleted unpaid order with ID: %v", order.ID)
		}
	}
}
