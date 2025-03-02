package service

import (
	"context"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/pb/frontend/user_page"
	"github.com/gin-gonic/gin"
)

type RegisterMerchantService struct {
	ctx context.Context
}

func NewRegisterMerchantService(c context.Context) *RegisterMerchantService {
	return &RegisterMerchantService{ctx: c}
}

func (s *RegisterMerchantService) Run(req *user_page.RegisterMerchantReq) (map[string]interface{}, error) {
	userClient, conn, err := rpcclient.GetUserClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	_, err = userClient.RegisterMerchant(s.ctx, &pbuser.RegisterMerchantReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return gin.H{}, nil
}
