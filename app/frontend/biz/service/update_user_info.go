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
	userClient, conn, err := rpcclient.GetUserClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := userClient.UpdateUserInfo(s.ctx, &pbuser.UpdateUserInfoReq{
		UserId:   req.UserId,
		Name:     req.Name,
		Sign:     req.Sign,
		Gender:   req.Gender,
		Birthday: req.Birthday,
	})
	if err != nil {
		return nil, err
	}

	return gin.H{
		"id":   resp.UserId,
		"name": resp.Name,
		"sign": resp.Sign,
	}, nil
}
