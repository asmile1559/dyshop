package service

import (
	"context"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/pb/frontend/user_page"
	"github.com/gin-gonic/gin"
)

// GetUserInfoService 注册服务
type GetAccountInfoService struct {
	ctx context.Context
}

// NewGetUserInfoService 创建注册服务实例
func NewGetAccountInfoService(c context.Context) *GetAccountInfoService {
	return &GetAccountInfoService{ctx: c}
}

func (s *GetAccountInfoService) Run(req *user_page.GetAccountInfoReq) (map[string]interface{}, error) {
	userClient, conn, err := rpcclient.GetUserClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := userClient.GetAccountInfo(s.ctx, &pbuser.GetAccountInfoReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return gin.H{
		"Id":    resp.UserId,
		"Name":  resp.Name,
		"Sign":  resp.Sign,
		"Img":   resp.Url, // 暂时指定头像图片url
		"Role":  resp.Role,
		"Phone": resp.Phone,
		"Email": resp.Email,
	}, nil
}
