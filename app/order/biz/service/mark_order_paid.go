package service

import (
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"golang.org/x/net/context"
)

type MarkOrderPaidService struct {
	ctx context.Context
}

func NewMarkOrderPaidService(c context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: c}
}

func (s *MarkOrderPaidService) Run(req *pborder.MarkOrderPaidReq) (*pborder.MarkOrderPaidResp, error) {
	// TODO: finish your business code...
	//
	return &pborder.MarkOrderPaidResp{}, nil
}
