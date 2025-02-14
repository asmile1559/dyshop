package service

import (
	"context"
	"fmt"

	"github.com/asmile1559/dyshop/app/user/biz/dal/mysql"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/sirupsen/logrus"

)

// GetUserInfoService 注册服务
type GetUserInfoService struct {
	ctx context.Context
}

// NewGetUserInfoService 创建注册服务实例
func NewGetUserInfoService(c context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: c}
}

// Run 执行获取用户信息逻辑
func (s *GetUserInfoService) Run(req *pbuser.GetUserInfoReq) (*pbuser.GetUserInfoResp, error) {
	// 1. 查找用户
	user, err := mysql.GetUserByID(uint32(req.UserId))
	if err != nil {
		logrus.WithField("userid",req.UserId).WithError(err).Error("用户不存在")
		return nil, fmt.Errorf("用户不存在: %v", err)
	}
	logrus.Info("get user info success")
	
	// 2. 返回用户信息
	return &pbuser.GetUserInfoResp{
		UserId:  user.UserID,
		Email:   user.Email,
		Password: user.Password, 
	}, nil
}
