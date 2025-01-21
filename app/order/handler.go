package main

import (
	service "github.com/asmile1559/dyshop/app/order/biz/service"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"golang.org/x/net/context"
)

type OrderServiceServer struct {
	pborder.UnimplementedOrderServiceServer
}

func (s *OrderServiceServer) PlaceOrder(ctx context.Context, req *pborder.PlaceOrderReq) (*pborder.PlaceOrderResp, error) {
	resp, err := service.NewPlaceOrderService(ctx).Run(req)

	return resp, err
}
func (s *OrderServiceServer) ListOrder(ctx context.Context, req *pborder.ListOrderReq) (*pborder.ListOrderResp, error) {
	resp, err := service.NewListOrdersService(ctx).Run(req)

	return resp, err
}
func (s *OrderServiceServer) MarkOrderPaid(ctx context.Context, req *pborder.MarkOrderPaidReq) (*pborder.MarkOrderPaidResp, error) {
	resp, err := service.NewMarkOrderPaidService(ctx).Run(req)

	return resp, err
}
