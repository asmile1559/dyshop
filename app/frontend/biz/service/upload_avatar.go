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
	userClient, conn, err := rpcclient.GetUserClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := userClient.UploadAvatar(s.ctx, &pbuser.UploadAvatarReq{
		UserId:    req.UserId,
		Filename:  req.Filename,
		ImageData: req.ImageData,
	})
	if err != nil {
		return nil, err
	}

	return gin.H{
		"user_id": resp.UserId,
		"url":     resp.Url,
	}, nil
}
