package service

import (
	"context"
	"fmt"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/app/user/biz/dal/mysql"
	"github.com/asmile1559/dyshop/app/user/utils"
	"github.com/asmile1559/dyshop/app/user/biz/model"
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
	// 1. 验证密码是否匹配
	if req.Password != req.ConfirmPassword {
		return nil, fmt.Errorf("密码不匹配")
	}

	// 2. 检查邮箱是否已注册
	existingUser, err := mysql.GetUserByEmail(req.Email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("邮箱已被注册")
	}

	// 3. 密码加密
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("密码加密失败: %v", err)
	}

	// 4. 存储用户信息到数据库
	newUser := &model.User{
		Email:    req.Email,
		Password: hashedPassword,
	}
	err = mysql.CreateUser(newUser)
	if err != nil {
		return nil, fmt.Errorf("用户创建失败: %v", err)
	}

	// 5. 返回用户 ID
	return &pbuser.RegisterResp{UserId: uint32(newUser.UserID)}, nil
}
