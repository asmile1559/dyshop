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

func (s *UserServiceServer) UpdateUserInfo(ctx context.Context, req *pbuser.UpdateUserInfoReq) (*pbuser.UpdateUserInfoResp, error) {
	return service.NewUpdateUserInfoService(ctx).Run(req)
}

func (s *UserServiceServer) GetUserInfo(ctx context.Context, req *pbuser.GetUserInfoReq) (*pbuser.GetUserInfoResp, error) {	
	return service.NewGetUserInfoService(ctx).Run(req)
}

func (s *UserServiceServer) GetAccountInfo(ctx context.Context, req *pbuser.GetAccountInfoReq) (*pbuser.GetAccountInfoResp, error) {	
	return service.NewGetAccountInfoService(ctx).Run(req)
}

func (s *UserServiceServer) UpdateAccount(ctx context.Context, req *pbuser.UpdateAccountReq) (*pbuser.UpdateAccountResp, error) {	
	return service.NewUpdateAccountService(ctx).Run(req)
}

func (s *UserServiceServer) UploadAvatar(ctx context.Context, req *pbuser.UploadAvatarReq) (*pbuser.UploadAvatarResp, error) {	
	return service.NewUploadAvatarService(ctx).Run(req)
}

func (s *UserServiceServer) RegisterMerchant(ctx context.Context, req *pbuser.RegisterMerchantReq) (*pbuser.RegisterMerchantResp, error) {	
	return service.NewRegisterMerchantService(ctx).Run(req)
}

func (s *UserServiceServer) DeleteUser(ctx context.Context, req *pbuser.DeleteUserReq) (*pbuser.DeleteUserResp, error) {	
	return service.NewDeleteUserService(ctx).Run(req)
}