package service

import (
	"context"

	"github.com/asmile1559/dyshop/app/cart/biz/dal"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
)

type EmptyCartService struct {
	ctx context.Context
}

func NewEmptyCartService(c context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: c}
}

func (s *EmptyCartService) Run(req *pbcart.EmptyCartReq) (*pbcart.EmptyCartResp, error) {
	// 清空购物车
	dal.ClearCart(req.UserId)

	return &pbcart.EmptyCartResp{}, nil
}
