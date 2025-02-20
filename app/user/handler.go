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
	return service.NewRegisterService(ctx).Run(req)
}

func (s *UserServiceServer) Login(ctx context.Context, req *pbuser.LoginReq) (*pbuser.LoginResp, error) {
	return service.NewLoginService(ctx).Run(req)
}

func (s *UserServiceServer) Logout(ctx context.Context, req *pbuser.LogoutReq) (*pbuser.LogoutResp, error) {
	return service.NewLogoutService(ctx).Run(req)
}
func (s *UserServiceServer) UpdateUser(ctx context.Context, req *pbuser.UpdateUserReq) (*pbuser.UpdateUserResp, error) {
	return service.NewUpdateUserService(ctx).Run(req)
}

func (s *UserServiceServer) GetUserInfo(ctx context.Context, req *pbuser.GetUserInfoReq) (*pbuser.GetUserInfoResp, error) {	
	return service.NewGetUserInfoService(ctx).Run(req)
}

func (s *UserServiceServer) DeleteUser(ctx context.Context, req *pbuser.DeleteUserReq) (*pbuser.DeleteUserResp, error) {	
	return service.NewDeleteUserService(ctx).Run(req)
}