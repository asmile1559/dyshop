package service

import (
	"context"

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
    
	resp, err := rpcclient.UserClient.GetUserInfo(s.ctx, &pbuser.GetUserInfoReq{
		UserId: req.GetUserId(),
	})  
    if err != nil {
        return nil, err
    }

    return gin.H{
		"resp": resp,
	}, nil

}