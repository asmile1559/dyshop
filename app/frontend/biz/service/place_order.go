package service

import (
	"context"
	"errors"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
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

	id, ok := s.ctx.Value("user_id").(uint32)
	if ok == false {
		return nil, errors.New("expect user id")
	}

	reqAddr := req.GetAddress()
	orderItems := make([]*pborder.OrderItem, 0)
	for _, item := range req.OrderItems {
		orderItems = append(orderItems, &pborder.OrderItem{
			Item: &pbcart.CartItem{
				ProductId: item.GetProductId(),
				Quantity:  item.GetQuantity(),
			},
			Cost: item.GetCost(),
		})
	}

	resp, err := rpcclient.OrderClient.PlaceOrder(s.ctx, &pborder.PlaceOrderReq{
		UserId:       id,
		UserCurrency: req.GetUserCurrency(),
		Address: &pborder.Address{
			StreetAddress: reqAddr.GetStreetAddress(),
			City:          reqAddr.GetCity(),
			State:         reqAddr.GetState(),
			Country:       reqAddr.GetCountry(),
			ZipCode:       reqAddr.GetZipCode(),
		},
		Email:      req.GetEmail(),
		OrderItems: orderItems,
	})

	if err != nil {
		return nil, err
	}

	return gin.H{
		"order_id": resp.Order.GetOrderId(),
	}, nil

	//return gin.H{
	//	"status": "place_order ok",
	//}, nil
}
