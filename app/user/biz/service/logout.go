package service

import (
	"context"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
)

// LogoutService 登录服务
type LogoutService struct {
	ctx context.Context
}

// NewLogoutService 创建登录服务实例
func NewLogoutService(c context.Context) *LogoutService {
	return &LogoutService{ctx: c}
}

// Run 执行登录逻辑
func (s *LogoutService) Run(req *pbuser.LogoutReq) (*pbuser.LogoutResp, error){
	// code
	return &pbuser.LogoutResp{
        Success: true,
    }, nil
}