package service

import (
	"context"
	"github.com/asmile1559/dyshop/pb/frontend/user_page"
	"github.com/gin-gonic/gin"
)

//rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
//pbbackend "github.com/asmile1559/dyshop/pb/backend/user"

type RegisterService struct {
	ctx context.Context
}

func NewRegisterService(c context.Context) *RegisterService {
	return &RegisterService{ctx: c}
}

func (s *RegisterService) Run(req *user_page.RegisterReq) (map[string]interface{}, error) {
	//resp, err := rpcclient.UserClient.Register(s.ctx, &pbbackend.RegisterReq{
	//	Email:           req.Email,
	//	Password:        req.Password,
	//	ConfirmPassword: req.ConfirmPassword,
	//})
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &user_page.RegisterResp{Id: resp.UserId}, nil

	return gin.H{
		"status": "register ok!",
	}, nil
}
