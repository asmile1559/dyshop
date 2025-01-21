package service

import (
	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
	"golang.org/x/net/context"
)

type VerifyTokenService struct {
	ctx context.Context
}

func NewVerifyService(c context.Context) *VerifyTokenService {
	return &VerifyTokenService{ctx: c}
}

func (s *VerifyTokenService) Run(req *pbauth.VerifyTokenReq) (*pbauth.VerifyResp, error) {
	// TODO: finish your business code...
	//
	return &pbauth.VerifyResp{Res: true}, nil
}
