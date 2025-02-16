package service

import (
	"context"
	"fmt"

	"github.com/asmile1559/dyshop/app/user/biz/dal/mysql"
	"github.com/asmile1559/dyshop/app/user/biz/model"
	"github.com/asmile1559/dyshop/app/user/utils/bcrypt"
	"github.com/asmile1559/dyshop/app/user/utils/snowflake"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/sirupsen/logrus"
)

// RegisterService 注册服务
type RegisterService struct {
	ctx context.Context
}

// NewRegisterService 创建注册服务实例
func NewRegisterService(c context.Context) *RegisterService {
	return &RegisterService{ctx: c}
}

// Run 执行注册逻辑
func (s *RegisterService) Run(req *pbuser.RegisterReq) (*pbuser.RegisterResp, error) {
	// 1. 检查邮箱是否已注册
	existingUser, err := mysql.GetUserByEmail(req.Email)
	if err == nil && existingUser != nil {
		logrus.WithField("register_email",req.Email).Error("邮箱已被注册")
		return nil, fmt.Errorf("邮箱已被注册")
	}

	// 2. 密码加密
	hashedPassword, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		logrus.WithError(err).Error("密码加密失败")
		return nil, fmt.Errorf("密码加密失败: %v", err)
	}

	// 3. 存储用户信息到数据库
	userID := snowflake.GenID()
	newUser := &model.User{
		UserID: userID,
		Email:    req.Email,
		Password: hashedPassword,
	}
	err = mysql.CreateUser(newUser)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"email":    req.Email,
			"password": hashedPassword,
		}).WithError(err).Error("用户创建失败")
		return nil, fmt.Errorf("用户创建失败: %v", err)
	}
	logrus.Info("register success")
	// 4. 返回用户 ID
	return &pbuser.RegisterResp{UserId: newUser.UserID}, nil
}
