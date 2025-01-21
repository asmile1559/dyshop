package main

import (
	"context"
	service "github.com/asmile1559/dyshop/app/user/biz/service"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
)

type UserServiceServer struct {
	pbuser.UnimplementedUserServiceServer
}

func (s *UserServiceServer) Register(ctx context.Context, req *pbuser.RegisterReq) (*pbuser.RegisterResp, error) {
	resp, err := service.NewRegisterService(ctx).Run(req)

	return resp, err
}
func (s *UserServiceServer) Login(ctx context.Context, req *pbuser.LoginReq) (*pbuser.LoginResp, error) {
	resp, err := service.NewLoginService(ctx).Run(req)

	return resp, err
}
