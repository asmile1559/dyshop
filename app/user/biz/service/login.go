package service

import (
	"context"
	"fmt"

	"github.com/asmile1559/dyshop/app/user/biz/dal/mysql"
	"github.com/asmile1559/dyshop/app/user/utils"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/sirupsen/logrus"
	"github.com/asmile1559/dyshop/utils/jwt"
)

// LoginService 登录服务
type LoginService struct {
	ctx context.Context
}

// NewLoginService 创建登录服务实例
func NewLoginService(c context.Context) *LoginService {
	return &LoginService{ctx: c}
}

// Run 执行登录逻辑
func (s *LoginService) Run(req *pbuser.LoginReq) (*pbuser.LoginResp, error) {
	// 1. 查询用户
	user, err := mysql.GetUserByEmail(req.Email)
	if err != nil {
		// 用户不存在
		return nil, fmt.Errorf("用户不存在: %v", err)
	}
	
	// 2. 验证密码
	if !utils.VerifyPassword(user.Password, req.Password) {
		// 密码不匹配
		return nil, fmt.Errorf("密码错误")
	}

	// 3. 生成令牌
	token, err := jwt.GenerateJWT(user.UserID)
	
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	
	// 4. 返回令牌和用户 ID
	return &pbuser.LoginResp{
		UserId: uint32(user.UserID),
		Token:  token, // 返回生成的令牌
	}, nil
}
