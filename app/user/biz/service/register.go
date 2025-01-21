package service

import (
	"context"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
)

type RegisterService struct {
	ctx context.Context
}

func NewRegisterService(c context.Context) *RegisterService {
	return &RegisterService{ctx: c}
}

func (s *RegisterService) Run(req *pbuser.RegisterReq) (*pbuser.RegisterResp, error) {
	// TODO: finish your business code...
	//
	return &pbuser.RegisterResp{UserId: 1}, nil
}
