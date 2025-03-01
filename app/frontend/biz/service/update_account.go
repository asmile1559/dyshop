package service

import (
	"context"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/pb/frontend/user_page"
	"github.com/gin-gonic/gin"
)

type UpdateAccountService struct {
	ctx context.Context
}

func NewUpdateAccountService(c context.Context) *UpdateAccountService {
	return &UpdateAccountService{ctx: c}
}

func (s *UpdateAccountService) Run(req *user_page.UpdateAccountReq) (map[string]interface{}, error) {
	// 调用 user 微服务的更新接口
	resp, err := rpcclient.UserClient.UpdateAccount(s.ctx, &pbuser.UpdateAccountReq{
		UserId:          req.UserId,
		Phone:           req.Phone,
		Email:           req.Email,
		Password:        req.Password,
		NewPassword:     req.NewPassword,
		ConfirmPassword: req.ConfirmPassword,
	})
	if err != nil {
		return nil, err
	}

	// 如果更新成功，返回响应
	return gin.H{
		"id":   resp.UserId,
		"phone": resp.Phone,
		"email": resp.Email,
	}, nil

}
