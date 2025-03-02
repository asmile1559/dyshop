package service

import (
	"context"
	"fmt"

	"github.com/asmile1559/dyshop/app/user/biz/dal/mysql"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	//rpcclient "github.com/asmile1559/dyshop/app/user/rpc"
	"github.com/sirupsen/logrus"
	
)

// RegisterMerchantService 注册商户服务
type RegisterMerchantService struct {
	ctx context.Context
}

// NewRegisterMerchantService 创建注册商户实例
func NewRegisterMerchantService(c context.Context) *RegisterMerchantService {
	return &RegisterMerchantService{ctx: c}
}

// Run 执行注册商户逻辑
func (s *RegisterMerchantService) Run(req *pbuser.RegisterMerchantReq) (*pbuser.RegisterMerchantResp, error){
	// 1. 查询用户是否存在
	user, err := mysql.GetUserByID(req.UserId)
	if err != nil {
		logrus.WithField("userId",req.UserId).WithError(err).Error("用户不存在")
		return nil, fmt.Errorf("用户不存在或查询失败: %v", err)
	}

	// 2. 检查用户是否已是商户
	if user.Role == "merchant" {
		logrus.WithField("user_id", req.UserId).Error("用户已经是商户")
		return nil, fmt.Errorf("用户已经是商户")
	}

	// 3. 更新用户角色为商户
	if err := mysql.SetUserAsMerchant(req.UserId); err != nil {
		logrus.WithError(err).Error("更新用户角色失败")
		return nil, fmt.Errorf("更新用户角色失败: %v", err)
	}

	// 4. 返回成功响应
	logrus.Info("register merchant success")
	return &pbuser.RegisterMerchantResp{},nil
}