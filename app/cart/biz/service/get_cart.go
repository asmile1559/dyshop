package service

import (
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
	"golang.org/x/net/context"
)

type GetCartService struct {
	ctx context.Context
}

func NewGetCartService(c context.Context) *GetCartService {
	return &GetCartService{ctx: c}
}

func (s *GetCartService) Run(req *pbcart.GetCartReq) (*pbcart.GetCartResp, error) {
	return &pbcart.GetCartResp{Cart: &pbcart.Cart{
		UserId: 1,
		Items: []*pbcart.CartItem{
			{ProductId: 1, Quantity: 100},
		},
	}}, nil
}
