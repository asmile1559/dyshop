package service

import (
	"context"
	"fmt"

	"github.com/asmile1559/dyshop/app/user/biz/dal/mysql"
	rpcclient "github.com/asmile1559/dyshop/app/user/rpc"
	"github.com/asmile1559/dyshop/app/user/utils/bcrypt"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/sirupsen/logrus"

	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
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
		logrus.WithField("login_email", req.Email).WithError(err).Error("用户不存在")
		return nil, fmt.Errorf("用户不存在: %v", err)
	}

	// 2. 验证密码
	if !bcrypt.VerifyPassword(user.Password, req.Password) {
		// 密码不匹配
		logrus.Error("密码错误")
		return nil, fmt.Errorf("密码错误")
	}

	// 3. 生成令牌
	// 调用auth.DeliverTokenByRPC
	authClient, conn, err := rpcclient.GetAuthClient()
	if err != nil {
		logrus.WithError(err).Error("rpcclient.GetAuthClient fail")
		return nil, err
	}
	defer conn.Close()
	resp, err := authClient.DeliverTokenByRPC(s.ctx, &pbauth.DeliverTokenReq{UserId: user.UserID})
	if err != nil {
		logrus.WithError(err).Error("AuthClient.DeliverTokenByRPC fail")
		return nil, err
	}
	logrus.Info("login success")
	// 4. 返回令牌和用户 ID
	return &pbuser.LoginResp{
		Token: resp.Token, // 返回生成的令牌
	}, nil
}
