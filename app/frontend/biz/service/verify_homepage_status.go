package service

import (
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/pb/frontend/home_page"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type VerifyHomepageStatusService struct {
	ctx context.Context
}

func NewVerifyHomepageStatus(c context.Context) *VerifyHomepageStatusService {
	return &VerifyHomepageStatusService{ctx: c}
}

func (s *VerifyHomepageStatusService) Run(req *home_page.VerifyHomepageStatusReq) gin.H {
	verifyTokenResp, err := rpcclient.AuthClient.VerifyTokenByRPC(s.ctx, &pbauth.VerifyTokenReq{
		Token:  req.GetToken(),
		Method: "POST",
		Uri:    "/",
	})

	if err != nil {
		logrus.Error(err)
		return nil
	}

	if !verifyTokenResp.GetRes() {
		return gin.H{
			"resp": gin.H{
				"ok": false,
			},
		}
	}
	userInfoResp, err := rpcclient.UserClient.GetUserInfo(s.ctx, &pbuser.GetUserInfoReq{UserId: verifyTokenResp.GetUserId()})
	if err != nil {
		logrus.Error(err)
		return nil
	}

	return gin.H{
		"resp": gin.H{
			"ok":   verifyTokenResp.GetRes(),
			"Id":   verifyTokenResp.GetUserId(),
			"Name": userInfoResp.GetName(),
			"Img":  userInfoResp.GetUrl(),
		},
	}
}
