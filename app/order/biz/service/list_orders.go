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
)

type ListOrdersService struct {
	ctx context.Context
	DB  *gorm.DB
}

func NewListOrdersService(c context.Context) *ListOrdersService {
	return &ListOrdersService{ctx: c, DB: db.DB}
}

func (s *ListOrdersService) Run(req *pborder.ListOrderReq) (*pborder.ListOrderResp, error) {
	var orders []model.Order
	if err := s.DB.Find(&orders).Error; err != nil {
		logrus.Error("Failed to fetch orders:", err)
		return nil, fmt.Errorf("failed to fetch orders: %v", err)
	}

	filteredOrders := make([]*pborder.Order, 0)
	for _, order := range orders {
		if uint32(order.UserId) == req.UserId { // 确保类型匹配
			filteredOrders = append(filteredOrders, &pborder.Order{
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
			})
		}
	}

	return &pborder.ListOrderResp{
		Orders: filteredOrders,
	}, nil
}
