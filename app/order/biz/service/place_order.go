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
	modelAddress := model.Address{
		AddressId:   strconv.FormatUint(orderID, 10), // 使用唯一订单ID作为地址ID
		Recipient:   address.GetRecipient(),
		Phone:       address.GetPhone(),
		Province:    address.GetProvince(),
		City:        address.GetCity(),
		District:    address.GetDistrict(),
		Street:      address.GetStreet(),
		FullAddress: address.GetFullAddress(),
	}

	products := make([]model.Product, len(req.GetProducts()))
	for i, item := range req.GetProducts() {
		spec := model.ProductSpec{
			Name:  item.GetProductSpec().GetName(),
			Price: item.GetProductSpec().GetPrice(),
		}
		products[i] = model.Product{
			ProductID:   uint64(item.GetProductId()),
			ProductImg:  item.GetProductImg(),
			ProductName: item.GetProductName(),
			ProductSpec: spec,
			Quantity:    int(item.GetQuantity()),
			Currency:    item.GetCurrency(),
			Postage:     item.GetPostage(),
		}
	}

	newOrder := model.Order{
		OrderID:         orderID,
		UserId:          req.GetUserId(),
		UserCurrency:    req.GetUserCurrency(),
		Address:         modelAddress,
		Email:           req.GetEmail(),
		CreatedAt:       time.Now(),
		OrderPrice:      req.GetOrderPrice(),
		OrderPostage:    req.GetOrderPostage(),
		OrderDiscount:   req.GetDiscount(),
		OrderFinalPrice: req.GetOrderFinalPrice(),
		Products:        products,
	}

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		// 创建订单
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
		OrderId: strconv.FormatUint(orderID, 10), // 使用 strconv 包将 uint64 转换为字符串
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
		logrus.Infof("Found unpaid order with ID: %v", order.OrderID)
	}

	for _, order := range unpaidOrders { // 删除未支付且超时的订单
		if err := s.DB.WithContext(ctx).Delete(&order).Error; err != nil {
			logrus.Error("Failed to delete unpaid order:", err)
		} else {
			logrus.Infof("Deleted unpaid order with ID: %v", order.OrderID)
		}
	}
}
