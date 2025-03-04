package service

import (
	"context"
	"errors"
	"time"

	"github.com/asmile1559/dyshop/app/checkout/biz/dal"
	"github.com/asmile1559/dyshop/app/checkout/biz/model"
	pbcheckout "github.com/asmile1559/dyshop/pb/backend/checkout"
)

// GetOrderWithItemsService 结算服务
type GetOrderWithItemsService struct {
	ctx context.Context
}

// NewGetOrderWithItemsService 创建结算服务实例
func NewGetOrderWithItemsService(ctx context.Context) *GetOrderWithItemsService {
	return &GetOrderWithItemsService{ctx: ctx}
}

// 查询订单及其商品
func (s *GetOrderWithItemsService) Run(req *pbcheckout.GetOrderReq) (*pbcheckout.GetOrderResp, error) {
	// 查询订单基本信息
	var order model.OrderRecord
	err := dal.DB.Where("order_id = ?", req.OrderId).First(&order).Error
	if err != nil {
		return nil, errors.New("订单不存在")
	}

	// 查询订单商品
	var items []model.OrderItem
	err = dal.DB.Where("order_id = ?", req.OrderId).Find(&items).Error
	if err != nil {
		return nil, errors.New("订单商品查询失败")
	}

	// 转换为 proto 格式
	orderProto := &pbcheckout.OrderRecord{
		OrderId:      order.OrderID,
		UserId:       order.UserID,
		TransactionId: order.TransactionID,
		Recipient:    order.Recipient,
		Phone:        order.Phone,
		Province:     order.Province,
		City:         order.City,
		District:     order.District,
		Street:       order.Street,
		FullAddress:  order.FullAddress,
		TotalQuantity: int32(order.TotalQuantity),
		TotalPrice:   order.TotalPrice,
		Postage:      order.Postage,
		FinalPrice:   order.FinalPrice,
		CreatedAt:    order.CreatedAt.Format(time.RFC3339),
	}

	var itemsProto []*pbcheckout.OrderItem
	for _, item := range items {
		itemsProto = append(itemsProto, &pbcheckout.OrderItem{
			ProductId:   item.ProductID,
			ProductImg:  item.ProductImg,
			ProductName: item.ProductName,
			SpecName:    item.SpecName,
			SpecPrice:   item.SpecPrice,
			Quantity:    int32(item.Quantity),
			Postage:     item.Postage,
			Currency:    item.Currency,
		})
	}

	return &pbcheckout.GetOrderResp{
		Order: orderProto,
		Items: itemsProto,
	}, nil
}
