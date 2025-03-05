package service

import (
	"context"
	"errors"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"github.com/asmile1559/dyshop/pb/frontend/order_page"
	"github.com/gin-gonic/gin"
)

type PlaceOrderService struct {
	ctx context.Context
}

func NewPlaceOrderService(c context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: c}
}

func (s *PlaceOrderService) Run(req *order_page.PlaceOrderReq) (map[string]interface{}, error) {
	id, ok := s.ctx.Value("user_id").(int64)
	if ok == false {
		return nil, errors.New("expect user id")
	}

	orderClient, conn, err := rpcclient.GetOrderClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := orderClient.PlaceOrder(s.ctx, &pborder.PlaceOrderReq{
		UserId:     id,
		AddressId:  req.GetAddressId(),
		ProductIds: req.GetProductIds(),
		Price:      req.GetPrice(),
	})

	if err != nil {
		return nil, err
	}
	return gin.H{
		"order_id": resp.OrderId,
	}, nil
}
