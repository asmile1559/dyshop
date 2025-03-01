package service

import (
	"context"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/pb/frontend/user_page"
	"github.com/gin-gonic/gin"
)

type RegisterService struct {
	ctx context.Context
}

func NewRegisterService(c context.Context) *RegisterService {
	return &RegisterService{ctx: c}
}

func (s *RegisterService) Run(req *user_page.RegisterReq) (map[string]interface{}, error) {
	resp, err := rpcclient.UserClient.Register(s.ctx, &pbuser.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	})

	if err != nil {
		return nil, err
	}

	return gin.H{
		"resp": resp,
	}, nil
}
