package main

import (
	service "github.com/asmile1559/dyshop/app/auth/biz/service"
	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
	"golang.org/x/net/context"
)

type AuthServiceServer struct {
	pbauth.UnimplementedAuthServiceServer
}

func (s *AuthServiceServer) DeliverTokenByRPC(ctx context.Context, req *pbauth.DeliverTokenReq) (*pbauth.DeliveryResp, error) {

	resp, err := service.NewDeliverService(ctx).Run(req)
	
	return resp, err
}
func (s *AuthServiceServer) VerifyTokenByRPC(ctx context.Context, req *pbauth.VerifyTokenReq) (*pbauth.VerifyResp, error) {

	resp, err := service.NewVerifyService(ctx).Run(req)

	return resp, err
}
