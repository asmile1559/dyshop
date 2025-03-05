package main

import (
	service "github.com/asmile1559/dyshop/app/checkout/biz/service"
	pbcheckout "github.com/asmile1559/dyshop/pb/backend/checkout"
	"golang.org/x/net/context"
)

type CheckoutServiceServer struct {
	pbcheckout.UnimplementedCheckoutServiceServer
}

func (s *CheckoutServiceServer) Checkout(ctx context.Context, req *pbcheckout.CheckoutReq) (*pbcheckout.CheckoutResp, error) {

	resp, err := service.NewCheckoutService(ctx).Run(req)

	return resp, err
}

func (s *CheckoutServiceServer) GetOrderWithItems(ctx context.Context, req *pbcheckout.GetOrderReq) (*pbcheckout.GetOrderResp, error) {

	resp, err := service.NewGetOrderWithItemsService(ctx).Run(req)

	return resp, err
}
