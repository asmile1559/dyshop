package service

import (
	"context"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/pb/frontend/user_page"
	"github.com/gin-gonic/gin"
)

type UploadAvatarService struct {
	ctx context.Context
}

func NewUploadAvatarService(c context.Context) *UploadAvatarService {
	return &UploadAvatarService{ctx: c}
}

func (s *UploadAvatarService) Run(req *user_page.UploadAvatarReq) (map[string]interface{}, error) {
	// 调用 user 微服务的更新接口
	resp, err := rpcclient.UserClient.UploadAvatar(s.ctx, &pbuser.UploadAvatarReq{
		UserId: req.UserId,
		Url:    req.Url,
	})
	if err != nil {
		return nil, err
	}

	// 如果更新成功，返回响应
	return gin.H{
		"user_id": resp.UserId,
		"url":     resp.Url,
	}, nil

}
