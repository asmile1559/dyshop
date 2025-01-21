package service

import (
	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
	"golang.org/x/net/context"
)

type DeliverService struct {
	ctx context.Context
}

func NewDeliverService(c context.Context) *DeliverService {
	return &DeliverService{ctx: c}
}

func (s *DeliverService) Run(req *pbauth.DeliverTokenReq) (*pbauth.DeliveryResp, error) {
	// TODO: finish your business code...
	//
	return &pbauth.DeliveryResp{Token: "deliver token ok!"}, nil
}
