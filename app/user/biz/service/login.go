package service

import (
	"context"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
)

type LoginService struct {
	ctx context.Context
}

func NewLoginService(c context.Context) *LoginService {
	return &LoginService{ctx: c}
}

func (s *LoginService) Run(req *pbuser.LoginReq) (*pbuser.LoginResp, error) {
	// TODO: finish your business code...
	//
	return &pbuser.LoginResp{UserId: 1}, nil
}
