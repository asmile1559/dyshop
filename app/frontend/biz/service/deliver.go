package service

import (
	"context"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
	"github.com/asmile1559/dyshop/pb/frontend/auth_page"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DeliverTokenService struct {
	ctx context.Context
}

func NewDeliverTokenService(c context.Context) *DeliverTokenService {
	return &DeliverTokenService{ctx: c}
}

func (s *DeliverTokenService) Run(req *auth_page.DeliverTokenReq) (map[string]interface{}, error) {
	authClient, conn, err := rpcclient.GetAuthClient()
	if err != nil {
		logrus.WithError(err).Debug("GetAuthClient err")
		return nil, err
	}
	defer conn.Close()
	resp, err := authClient.DeliverTokenByRPC(s.ctx, &pbauth.DeliverTokenReq{UserId: req.UserId})
	if err != nil {
		return nil, err
	}

	return gin.H{
		"resp": resp,
	}, nil
}
