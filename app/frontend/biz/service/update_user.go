package service

import (
	"context"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/pb/frontend/user_page"
	"github.com/gin-gonic/gin"
)

type UpdateUserService struct {
	ctx context.Context
}

func NewUpdateUserService(c context.Context) *UpdateUserService {
	return &UpdateUserService{ctx: c}
}

func (s *UpdateUserService) Run(req *user_page.UpdateUserReq) (map[string]interface{}, error) {
	// 调用 user 微服务的更新接口，发送 email 和 password（如果提供）
	updateResp, err := rpcclient.UserClient.UpdateUser(s.ctx, &pbuser.UpdateUserReq{
		UserId:   req.UserId,
		Email:    req.Email,    // 如果更新邮箱，传入新邮箱
		Password: req.Password, // 如果更新密码，传入新密码
	})
	if err != nil {
		return nil, err
	}

	// 如果更新成功，返回响应
	return gin.H{
		"resp": updateResp,
	}, nil

}