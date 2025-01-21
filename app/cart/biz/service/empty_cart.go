package service

import (
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
	"golang.org/x/net/context"
)

type EmptyService struct {
	ctx context.Context
}

func NewEmptyService(c context.Context) *EmptyService {
	return &EmptyService{ctx: c}
}

func (s *EmptyService) Run(req *pbcart.EmptyCartReq) (*pbcart.EmptyCartResp, error) {
	// TODO: finish your business code...
	//
	return &pbcart.EmptyCartResp{}, nil
}
