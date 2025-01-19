package service

import (
	"context"
	"github.com/dyshop/pb/frontend/user_page"
)

type LoginService struct {
	ctx context.Context
}

func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

func (s *LoginService) Run(req *user_page.LoginReq) (resp *user_page.LoginResp, err error) {
	//loginResp, err := rpcclient.UserClient.Login(s.ctx, &backenduser.LoginReq{
	//	Email:    req.Email,
	//	Password: req.Password,
	//})
	//if err != nil {
	//	return nil, err
	//}

	return &user_page.LoginResp{
		Id: 1001,
	}, nil
}
