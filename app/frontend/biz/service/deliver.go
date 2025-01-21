package service

import (
	"context"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
	"github.com/asmile1559/dyshop/pb/frontend/auth_page"
	"github.com/gin-gonic/gin"
)

type DeliverTokenService struct {
	ctx context.Context
}

func NewDeliverTokenService(c context.Context) *DeliverTokenService {
	return &DeliverTokenService{ctx: c}
}

func (s *DeliverTokenService) Run(req *auth_page.DeliverTokenReq) (map[string]interface{}, error) {
	resp, err := rpcclient.AuthClient.DeliverTokenByRPC(s.ctx, &pbauth.DeliverTokenReq{UserId: req.UserId})
	if err != nil {
		return nil, err
	}

	return gin.H{
		"resp": resp,
	}, nil

	//return gin.H{
	//	"status": "deliver token ok",
	//}, nil
}
