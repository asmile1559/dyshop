package service

import (
	"context"
	"time"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/pb/frontend/user_page"
	"github.com/gin-gonic/gin"
)

// GetUserInfoService 注册服务
type GetUserInfoService struct {
	ctx context.Context
}

// NewGetUserInfoService 创建注册服务实例
func NewGetUserInfoService(c context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: c}
}

func (s *GetUserInfoService) Run(req *user_page.GetUserInfoReq) (map[string]interface{}, error) {
	userClient, conn, err := rpcclient.GetUserClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := userClient.GetUserInfo(s.ctx, &pbuser.GetUserInfoReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	birthday, _ := time.Parse("2006年1月2日", resp.Birthday)
	return gin.H{
		"Id":       resp.UserId,
		"Name":     resp.Name,
		"Sign":     resp.Sign,
		"Img":      resp.Url,
		"Role":     resp.Role,
		"Gender":   resp.Gender,
		"Birthday": birthday,
	}, nil
}
