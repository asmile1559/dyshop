package service

import (
	"context"
	"github.com/asmile1559/dyshop/pb/frontend/auth_page"
	"github.com/gin-gonic/gin"
)

//rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
//pbbackend "github.com/asmile1559/dyshop/pb/backend/auth"

type VerifyTokenService struct {
	ctx context.Context
}

func NewVerifyService(c context.Context) *VerifyTokenService {
	return &VerifyTokenService{ctx: c}
}

func (s *VerifyTokenService) Run(req *auth_page.VerifyTokenReq) (map[string]interface{}, error) {
	//resp, err := rpcclient.AuthClient.VerifyTokenByRPC(s.ctx, &pbbackend.VerifyTokenReq{Token: req.Token})
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &auth_page.VerifyResp{Res: resp.GetRes()}, nil

	return gin.H{
		"status": "verify ok!",
	}, nil
}
