package service

import (
	"context"
	pbpayment "github.com/asmile1559/dyshop/pb/backend/payment"
)

type ChargeService struct {
	ctx context.Context
}

func NewChargeService(c context.Context) *ChargeService {
	return &ChargeService{ctx: c}
}

func (s *ChargeService) Run(req *pbpayment.ChargeReq) (*pbpayment.ChargeResp, error) {
	// TODO: finish your business code...
	//
	return &pbpayment.ChargeResp{TransactionId: "123"}, nil
}
