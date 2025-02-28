package service

import (
	"context"
	"fmt"
	"github.com/asmile1559/dyshop/app/user/biz/dal/mysql"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/sirupsen/logrus"
	"github.com/asmile1559/dyshop/app/user/utils/bcrypt"

)

// UpdateAccountService 更新账户信息服务
type UpdateAccountService struct {
	ctx context.Context
}

// NewUpdateAccountService 创建更新账户信息服务实例
func NewUpdateAccountService(c context.Context) *UpdateAccountService {
	return &UpdateAccountService{ctx: c}
}

// Run 执行更新账户信息逻辑
func (s *UpdateAccountService) Run(req *pbuser.UpdateAccountReq) (*pbuser.UpdateAccountResp, error) {
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
	if req.NewPassword != "" {
		user.Password ,err = bcrypt.HashPassword(req.NewPassword)  // 记得对密码进行加密存储
		if err != nil {
			logrus.WithError(err).Error("密码加密失败")
			return nil, fmt.Errorf("密码加密失败: %v", err)
		}
	}
	user.Phone=req.Phone;
	
	// 4. 更新用户信息到数据库
	err = mysql.UpdateUser(user)
	if err != nil {
		logrus.WithError(err).Error("更新用户信息失败")
		return nil, fmt.Errorf("更新用户信息失败: %v", err)
	}
	logrus.Info("update account info success")
	
	// 5. 返回成功响应
	return &pbuser.UpdateAccountResp{
		UserId: user.UserID,
		Phone:  user.Phone,
		Email:  user.Email,
	}, nil
}
