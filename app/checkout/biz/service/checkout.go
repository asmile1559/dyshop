package service

import (
	pbcheckout "github.com/asmile1559/dyshop/pb/backend/checkout"
	"golang.org/x/net/context"
)

type CheckoutService struct {
	ctx context.Context
}

func NewCheckoutService(c context.Context) *CheckoutService {
	return &CheckoutService{ctx: c}
}

func (s *CheckoutService) Run(req *pbcheckout.CheckoutReq) (*pbcheckout.CheckoutResp, error) {
	// TODO: finish your business code...
	//
	return &pbcheckout.CheckoutResp{
		OrderId:       "123",
		TransactionId: "123",
	}, nil
}
