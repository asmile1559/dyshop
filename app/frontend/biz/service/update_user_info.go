package service

import (
	"context"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/pb/frontend/user_page"
	"github.com/gin-gonic/gin"
)

type UpdateUserInfoService struct {
	ctx context.Context
}

func NewUpdateUserInfoService(c context.Context) *UpdateUserInfoService {
	return &UpdateUserInfoService{ctx: c}
}

func (s *UpdateUserInfoService) Run(req *user_page.UpdateUserInfoReq) (map[string]interface{}, error) {
	// 调用 user 微服务的更新接口
	resp, err := rpcclient.UserClient.UpdateUserInfo(s.ctx, &pbuser.UpdateUserInfoReq{
		UserId:   req.UserId,
		Name:     req.Name,
		Sign:     req.Sign,
		Gender:   req.Gender,
		Birthday: req.Birthday,
	})
	if err != nil {
		return nil, err
	}

	// 如果更新成功，返回响应
	return gin.H{
		"id":   resp.UserId,
		"name": resp.Name,
		"sign": resp.Sign,
	}, nil

}
