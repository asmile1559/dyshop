package service

import (
	"context"
	"fmt"

	"github.com/asmile1559/dyshop/app/user/biz/dal/mysql"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/sirupsen/logrus"

)

// GetUserInfoService 获取用户信息服务
type GetUserInfoService struct {
	ctx context.Context
}

// NewGetUserInfoService 创建获取用户信息服务实例
func NewGetUserInfoService(c context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: c}
}

// Run 执行获取用户信息逻辑
func (s *GetUserInfoService) Run(req *pbuser.GetUserInfoReq) (*pbuser.GetUserInfoResp, error) {
	// 1. 查找用户
	user, err := mysql.GetUserByID(req.UserId)
	if err != nil {
		logrus.WithField("userid",req.UserId).WithError(err).Error("用户不存在")
		return nil, fmt.Errorf("用户不存在: %v", err)
	}
	logrus.Info("get user info success")
	
	// 2. 返回用户信息
	return &pbuser.GetUserInfoResp{
		UserId:   user.UserID,
		Name:     user.Name,
		Sign:     user.Sign,
		Url:      "http://localhost:12167"+user.Url,
		Role:     []string{user.Role},
		Gender:   user.Gender,
		Birthday: user.Birthday.Format("2006年1月2日"),
	}, nil
}
