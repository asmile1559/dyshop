package service

import (
	"context"
	"fmt"
	"os"

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
	
	userID := req.GetUserId()
    filename := req.GetFilename()
    imageData := req.GetImageData()

    if userID == 0 || filename == "" || len(imageData) == 0 {
		logrus.Error("userID, filename 或 imageData 为空")
        return nil, fmt.Errorf("userID, filename 或 imageData 为空")
    }

    // 创建存储目录
    fileDir := fmt.Sprintf("/static/src/user/%d/", userID)
    saveDir := "." + fileDir

    if _, err := os.Stat(saveDir); os.IsNotExist(err) {
        err = os.MkdirAll(saveDir, 0755)
        if err != nil {
			logrus.WithError(err).Error("创建目录失败")
            return nil, fmt.Errorf("创建目录失败: %w", err)
        }
    }

    // 保存文件
	filePath := fileDir + filename
	savePath := saveDir + filename
    err := os.WriteFile(savePath, imageData, 0644)
    if err != nil {
		logrus.WithError(err).Error("保存文件失败")
        return nil, fmt.Errorf("保存文件失败: %w", err)
    }

	if err := mysql.UpdateUserAvatar(req.UserId, filePath); err != nil {
		logrus.WithError(err).Error("更新数据库头像 URL 失败")
		return nil, fmt.Errorf("更新数据库头像 URL 失败")
	}

	// 5. 返回成功响应
	logrus.Info("update avatar success")
	return &pbuser.UploadAvatarResp{
		UserId: userID,
		Url:    "http://localhost:12167"+filePath,
	},nil
}