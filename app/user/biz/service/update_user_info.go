package service

import (
	"context"
	"fmt"
	"time"

	"github.com/asmile1559/dyshop/app/user/biz/dal/mysql"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/sirupsen/logrus"
)

// UpdateUserInfoService 更新用户信息服务
type UpdateUserInfoService struct {
	ctx context.Context
}

// NewUpdateUserInfoService 创建更新用户信息服务实例
func NewUpdateUserInfoService(c context.Context) *UpdateUserInfoService {
	return &UpdateUserInfoService{ctx: c}
}

// Run 执行更新用户信息逻辑
func (s *UpdateUserInfoService) Run(req *pbuser.UpdateUserInfoReq) (*pbuser.UpdateUserInfoResp, error) {
	// 1. 查询用户
	user, err := mysql.GetUserByID(req.UserId)
	if err != nil {
		// 用户不存在
		logrus.WithField("userId",req.UserId).WithError(err).Error("用户不存在")
		return nil, fmt.Errorf("用户不存在: %v", err)
	}

	// 2. 更新
	user.Name=req.Name;
	user.Sign=req.Sign;
	user.Gender=req.Gender;
	user.Birthday,_=time.Parse("2006年1月2日", req.Birthday);
	// 3. 更新用户信息到数据库
	err = mysql.UpdateUser(user)
	if err != nil {
		logrus.WithError(err).Error("更新用户信息失败")
		return nil, fmt.Errorf("更新用户信息失败: %v", err)
	}
	logrus.Info("update user info success")
	
	// 4. 返回成功响应
	return &pbuser.UpdateUserInfoResp{
		UserId: req.UserId,
		Name:   req.Name,
		Sign:   req.Sign,
	}, nil
}
