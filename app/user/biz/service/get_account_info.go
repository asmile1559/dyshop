package service

import (
	"context"
	"fmt"

	"github.com/asmile1559/dyshop/app/user/biz/dal/mysql"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/sirupsen/logrus"

)

// GetAccountInfoService 获取账户信息
type GetAccountInfoService struct {
	ctx context.Context
}

// NewGetAccountInfoService 创建获取账户信息实例
func NewGetAccountInfoService(c context.Context) *GetAccountInfoService {
	return &GetAccountInfoService{ctx: c}
}

// Run 执行获取账户信息逻辑
func (s *GetAccountInfoService) Run(req *pbuser.GetAccountInfoReq) (*pbuser.GetAccountInfoResp, error) {
	// 1. 查找用户
	user, err := mysql.GetUserByID(req.UserId)
	if err != nil {
		logrus.WithField("userid",req.UserId).WithError(err).Error("用户不存在")
		return nil, fmt.Errorf("用户不存在: %v", err)
	}
	logrus.Info("get account info success")
	
	// 2. 返回用户信息
	return &pbuser.GetAccountInfoResp{
		UserId: user.UserID,
		Name:   user.Name,
		Sign:   user.Sign,
		Role:   []string{user.Role},
		Phone:  user.Phone,
		Email:  user.Email,
		Url:    "http://localhost:12167"+user.Url,
	}, nil
}
