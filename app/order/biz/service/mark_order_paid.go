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
	var order model.Order
	if err := s.DB.First(&order, req.GetOrderId()).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("order not found")
		}
		logrus.Error("Failed to fetch order:", err)
		return nil, fmt.Errorf("failed to fetch order: %v", err)
	}

	// 更新订单状态
	order.Paid = true
	order.PaidAt = time.Now()

	if err := s.DB.Save(&order).Error; err != nil {
		logrus.Error("Failed to update order status:", err)
		return nil, fmt.Errorf("failed to update order status: %v", err)
	}

	// 构建响应
	respOrders := make([]*pborder.Order, 1)
	respOrders[0] = &pborder.Order{
		OrderId:      strconv.FormatUint(order.ID, 10),
		UserId:       uint32(order.UserId), // 确保类型匹配
		UserCurrency: order.UserCurrency,
		Address: &pborder.Address{
			StreetAddress: order.Address.StreetAddress,
			City:          order.Address.City,
			State:         order.Address.State,
			Country:       order.Address.Country,
			ZipCode:       order.Address.ZipCode,
		},
		Email:     order.Email,
		CreatedAt: int32(order.CreatedAt.Unix()), // 确保类型匹配
		OrderItems: func() []*pborder.OrderItem {
			items := make([]*pborder.OrderItem, len(order.OrderItems))
			for i, item := range order.OrderItems {
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
