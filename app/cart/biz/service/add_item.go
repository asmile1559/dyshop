package service

import (
	"context"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
)

type AddItemService struct {
	ctx context.Context
}

func NewAddItemService(c context.Context) *AddItemService {
	return &AddItemService{ctx: c}
}

func (s *AddItemService) Run(req *pbcart.AddItemReq) (*pbcart.AddItemResp, error) {
	// TODO: finish your business code...
	//
	return &pbcart.AddItemResp{}, nil
}
