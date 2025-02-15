package service

import (
	"context"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/pb/frontend/user_page"
	"github.com/gin-gonic/gin"
)

// DeleteUserService 注册服务
type DeleteUserService struct {
	ctx context.Context
}

// NewDeleteUserService 创建注册服务实例
func NewDeleteUserService(c context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: c}
}

func (s *DeleteUserService) Run(req *user_page.DeleteUserReq) (map[string]interface{}, error) {
   
	resp, err := rpcclient.UserClient.DeleteUser(s.ctx, &pbuser.DeleteUserReq{
		UserId: req.UserId,
	})  
    if err != nil {
        return nil, err
    }

    return gin.H{
		"resp": resp,
	}, nil

}