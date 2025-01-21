package service

import (
	"context"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
)

type PlaceOrderService struct {
	ctx context.Context
}

func NewPlaceOrderService(c context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: c}
}

func (s *PlaceOrderService) Run(req *pborder.PlaceOrderReq) (*pborder.PlaceOrderResp, error) {
	// TODO: finish your business code...
	//
	return &pborder.PlaceOrderResp{
		Order: &pborder.OrderResult{OrderId: "1"},
	}, nil

}
