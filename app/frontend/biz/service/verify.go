package service

import (
	"context"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
	"github.com/asmile1559/dyshop/pb/frontend/auth_page"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type VerifyTokenService struct {
	ctx context.Context
}

func NewVerifyService(c context.Context) *VerifyTokenService {
	return &VerifyTokenService{ctx: c}
}

func (s *VerifyTokenService) Run(req *auth_page.VerifyTokenReq) (map[string]interface{}, error) {
	authClient, conn, err := rpcclient.GetAuthClient()
	if err != nil {
		logrus.WithError(err).Debug("GetAuthClient err")
		return nil, err
	}
	defer conn.Close()
	resp, err := authClient.VerifyTokenByRPC(s.ctx, &pbauth.VerifyTokenReq{Token: req.Token})
	if err != nil {
		return nil, err
	}

	return gin.H{
		"resp": resp,
	}, nil

	//return gin.H{
	//	"status": "verify ok!",
	//}, nil
}
