package service

import (
	"context"
	"fmt"

	"github.com/asmile1559/dyshop/app/user/biz/dal/mysql"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/sirupsen/logrus"
)

// UploadAvatarService 上传头像服务
type UploadAvatarService struct {
	ctx context.Context
}

// NewUploadAvatarService 创建上传头像服务实例
func NewUploadAvatarService(c context.Context) *UploadAvatarService {
	return &UploadAvatarService{ctx: c}
}

// Run 执行上传头像逻辑
func (s *UploadAvatarService) Run(req *pbuser.UploadAvatarReq) (*pbuser.UploadAvatarResp, error){
	
	if err := mysql.UpdateUserAvatar(req.UserId, req.Url); err != nil {
		logrus.WithError(err).Error("更新数据库头像 URL 失败")
		return nil, fmt.Errorf("更新数据库头像 URL 失败")
	}

	// 5. 返回成功响应
	logrus.Info("update avatar success")
	return &pbuser.UploadAvatarResp{
		UserId: req.UserId,
		Url:    req.Url,
	},nil
}

