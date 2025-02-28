package service

import (
	"context"
	"fmt"
	"github.com/asmile1559/dyshop/app/order/biz/model"
	"github.com/asmile1559/dyshop/app/order/utils/db"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type MarkOrderPaidService struct {
	ctx context.Context
	DB  *gorm.DB
}

func NewMarkOrderPaidService(c context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: c, DB: db.DB}
}

func (s *MarkOrderPaidService) Run(req *pborder.MarkOrderPaidReq) (*pborder.MarkOrderPaidResp, error) {
	var prePaidOrder model.PrePaidOrder
	if err := s.DB.Preload("Address").Preload("OrderItems").First(&prePaidOrder, req.GetOrderId()).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("order not found")
		}
		logrus.Error("Failed to fetch prepaid order:", err)
		return nil, fmt.Errorf("failed to fetch prepaid order: %v", err)
	}

	// 创建新的 Order 对象并复制数据
	newOrder := model.Order{
		ID:           prePaidOrder.ID,
		UserId:       prePaidOrder.UserId,
		UserCurrency: prePaidOrder.UserCurrency,
		Address:      model.Address(prePaidOrder.Address),
		Email:        prePaidOrder.Email,
		Paid:         true,
		OrderItems:   make([]model.OrderItem, len(prePaidOrder.OrderItems)),
		CreatedAt:    prePaidOrder.CreatedAt,
		UpdatedAt:    time.Now(),
		PaidAt:       time.Now(),
	}

	for i, item := range prePaidOrder.OrderItems {
		newOrder.OrderItems[i] = model.OrderItem(item)
	}

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		// 创建正式订单
		if err := tx.Create(&newOrder).Error; err != nil {
			logrus.Error("Failed to create order:", err)
			return err
		}

		// 删除预支付订单
		if err := tx.Delete(&prePaidOrder).Error; err != nil {
			logrus.Error("Failed to delete prepaid order:", err)
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	// 构建响应
	respOrders := make([]*pborder.Order, 1)
	respOrders[0] = &pborder.Order{
		OrderId:      strconv.FormatUint(newOrder.ID, 10),
		UserId:       uint32(newOrder.UserId), // 确保类型匹配
		UserCurrency: newOrder.UserCurrency,
		Address: &pborder.Address{
			StreetAddress: newOrder.Address.StreetAddress,
			City:          newOrder.Address.City,
			State:         newOrder.Address.State,
			Country:       newOrder.Address.Country,
			ZipCode:       newOrder.Address.ZipCode,
		},
		Email:     newOrder.Email,
		CreatedAt: int32(newOrder.CreatedAt.Unix()), // 确保类型匹配
		OrderItems: func() []*pborder.OrderItem {
			items := make([]*pborder.OrderItem, len(newOrder.OrderItems))
			for i, item := range newOrder.OrderItems {
				items[i] = &pborder.OrderItem{
					Item: &pbcart.CartItem{
						ProductId: uint32(item.ProductId),
						Quantity:  int32(item.Quantity),
					},
					Cost: float32(item.Cost), // 确保类型匹配
				}
			}
			return items
		}(),
	}

	return &pborder.MarkOrderPaidResp{
		Orders: respOrders,
	}, nil
}
