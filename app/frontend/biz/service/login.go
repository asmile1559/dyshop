package service

import (
	"context"
	"github.com/asmile1559/dyshop/pb/frontend/user_page"
	"github.com/gin-gonic/gin"
)

//rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
//pbbackend "github.com/asmile1559/dyshop/pb/backend/user"

type LoginService struct {
	ctx context.Context
}

func NewLoginService(c context.Context) *LoginService {
	return &LoginService{ctx: c}
}

func (s *LoginService) Run(req *user_page.LoginReq) (map[string]interface{}, error) {
	//loginResp, err := rpcclient.UserClient.Login(s.ctx, &pbbackend.LoginReq{
	//	Email:    req.Email,
	//	Password: req.Password,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//return gin.H{
	//	"resp": resp,
	//}, nil

	return gin.H{
		"status": "login ok!",
	}, nil
}
