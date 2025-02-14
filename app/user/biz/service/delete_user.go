package service

import (
	"context"
	"fmt"

	"github.com/asmile1559/dyshop/app/user/biz/dal/mysql"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/sirupsen/logrus"
)

// DeleteUserService 注册服务
type DeleteUserService struct {
	ctx context.Context
}

// NewDeleteUserService 创建注册服务实例
func NewDeleteUserService(c context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: c}
}

// Run 执行删除用户逻辑
func (s *DeleteUserService) Run(req *pbuser.DeleteUserReq) (*pbuser.DeleteUserResp, error) {
	// 1. 查找用户
    user, err := mysql.GetUserByID(uint32(req.UserId))
    if err != nil {
        logrus.WithField("userId",req.UserId).WithError(err).Error("用户不存在")
        return nil, fmt.Errorf("用户不存在: %v", err)
    }

    // 2. 删除用户
    err = mysql.DeleteUserByID(user.UserID)
    if err != nil {
        logrus.WithField("userId",req.UserId).Error("删除用户失败")
        return nil, fmt.Errorf("删除用户失败: %v", err)
    }
    logrus.Info("delete user success")
    
    return &pbuser.DeleteUserResp{
        Success: true,
    }, nil
}
