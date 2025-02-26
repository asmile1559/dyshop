package service

import (
	"context"
	"fmt"
	"github.com/asmile1559/dyshop/app/user/biz/dal/mysql"
	"github.com/asmile1559/dyshop/app/user/utils"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/sirupsen/logrus"

)

// UpdateUserService 用户更新服务
type UpdateUserService struct {
	ctx context.Context
}

// NewUpdateUserService 创建更新用户服务实例
func NewUpdateUserService(c context.Context) *UpdateUserService {
	return &UpdateUserService{ctx: c}
}

// Run 执行更新用户逻辑
func (s *UpdateUserService) Run(req *pbuser.UpdateUserReq) (*pbuser.UpdateUserResp, error) {
	// 1. 查询用户
	user, err := mysql.GetUserByID(req.UserId)
	if err != nil {
		// 用户不存在
		logrus.WithField("userId",req.UserId).WithError(err).Error("用户不存在")
		return nil, fmt.Errorf("用户不存在: %v", err)
	}

	// 2. 如果提供了新邮箱，检查邮箱是否已被使用
	if req.Email != "" && req.Email != user.Email {
		existingUser, err := mysql.GetUserByEmail(req.Email)
		if err == nil && existingUser != nil {
			// 如果邮箱已被其他用户使用，返回错误
			logrus.WithField("email",req.Email).Error("该邮箱已被使用")
			return nil, fmt.Errorf("该邮箱已被使用")
		}
		user.Email = req.Email  // 更新邮箱
	}

	// 3. 如果提供了新密码，更新密码
	if req.Password != "" {
		user.Password ,err = utils.HashPassword(req.Password)  // 记得对密码进行加密存储
		if err != nil {
			logrus.WithError(err).Error("密码加密失败")
			return nil, fmt.Errorf("密码加密失败: %v", err)
		}
	}

	// 4. 更新用户信息到数据库
	err = mysql.UpdateUser(user)
	if err != nil {
		logrus.WithError(err).Error("更新用户信息失败")
		return nil, fmt.Errorf("更新用户信息失败: %v", err)
	}
	logrus.Info("update user success")
	
	// 5. 返回成功响应
	return &pbuser.UpdateUserResp{
		Success: true,
	}, nil
}
