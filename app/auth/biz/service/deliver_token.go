package service

import (
	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
	"github.com/asmile1559/dyshop/utils/jwt"
	"golang.org/x/net/context"
)

type DeliverService struct {
	ctx context.Context
}

func NewDeliverService(c context.Context) *DeliverService {
	return &DeliverService{ctx: c}
}

func (s *DeliverService) Run(req *pbauth.DeliverTokenReq) (*pbauth.DeliveryResp, error) {

	token, err := jwt.GenerateJWT(req.UserId)
	if err != nil {
		return nil, err
	}

	return &pbauth.DeliveryResp{Token: token}, nil
}
