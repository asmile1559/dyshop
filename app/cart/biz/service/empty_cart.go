package service

import (
	"context"

	"github.com/asmile1559/dyshop/app/cart/biz/dal"
	"github.com/asmile1559/dyshop/app/cart/biz/model"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
)

type EmptyCartService struct {
	ctx context.Context
}

func NewEmptyCartService(c context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: c}
}

func (s *EmptyCartService) Run(req *pbcart.EmptyCartReq) (*pbcart.EmptyCartResp, error) {
	err := model.ClearCart(dal.DB, uint64(req.UserId))
	if err != nil {
		return nil, err
	}
	return &pbcart.EmptyCartResp{}, nil
}
