package service

import (
	"context"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/pb/frontend/user_page"
	"github.com/gin-gonic/gin"
)

type LoginService struct {
	ctx context.Context
}

func NewLoginService(c context.Context) *LoginService {
	return &LoginService{ctx: c}
}

func (s *LoginService) Run(req *user_page.LoginReq) (map[string]interface{}, error) {
	resp, err := rpcclient.UserClient.Login(s.ctx, &pbuser.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return gin.H{
		"token": resp.Token,
	}, nil
}
