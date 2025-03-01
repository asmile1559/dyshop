package main

import (
	"context"

	service "github.com/asmile1559/dyshop/app/cart/biz/service"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
)

type CartServiceServer struct {
	pbcart.UnimplementedCartServiceServer
}

// 实现了 proto 中的 AddItem 方法
func (s *CartServiceServer) AddItem(ctx context.Context, req *pbcart.AddItemReq) (*pbcart.AddItemResp, error) {
	return service.NewAddItemService(ctx).Run(req)
}

// 实现了 proto 中的 GetCart 方法
func (s *CartServiceServer) GetCart(ctx context.Context, req *pbcart.GetCartReq) (*pbcart.GetCartResp, error) {
	return service.NewGetCartService(ctx).Run(req)
}

// 实现了 proto 中的 EmptyCart 方法
func (s *CartServiceServer) EmptyCart(ctx context.Context, req *pbcart.EmptyCartReq) (*pbcart.EmptyCartResp, error) {
	return service.NewEmptyCartService(ctx).Run(req)
}

// 实现了 proto 中的 DeleteCart 方法
func (s *CartServiceServer) DeleteCart(ctx context.Context, req *pbcart.DeleteCartReq) (*pbcart.DeleteCartResp, error) {
	return service.NewDeleteCartService(ctx).Run(req)
}
