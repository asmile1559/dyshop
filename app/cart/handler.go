package main

import (
	service "github.com/asmile1559/dyshop/app/cart/biz/service"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
	"golang.org/x/net/context"
)

type CartServiceServer struct {
	pbcart.UnimplementedCartServiceServer
}

func (s *CartServiceServer) AddItem(ctx context.Context, req *pbcart.AddItemReq) (*pbcart.AddItemResp, error) {

	resp, err := service.NewAddItemService(ctx).Run(req)

	return resp, err
}
func (s *CartServiceServer) GetCart(ctx context.Context, req *pbcart.GetCartReq) (*pbcart.GetCartResp, error) {

	resp, err := service.NewGetCartService(ctx).Run(req)

	return resp, err
}
func (s *CartServiceServer) EmptyCart(ctx context.Context, req *pbcart.EmptyCartReq) (*pbcart.EmptyCartResp, error) {

	resp, err := service.NewEmptyService(ctx).Run(req)

	return resp, err
}
