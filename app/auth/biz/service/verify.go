package service

import (
	"github.com/asmile1559/dyshop/app/auth/utils/casbin"
	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
	"github.com/asmile1559/dyshop/utils/jwt"
	"golang.org/x/net/context"
	"net/http"
)

type VerifyTokenService struct {
	ctx context.Context
}

func NewVerifyService(c context.Context) *VerifyTokenService {
	return &VerifyTokenService{ctx: c}
}

func (s *VerifyTokenService) Run(req *pbauth.VerifyTokenReq) (*pbauth.VerifyResp, error) {

	resp := pbauth.VerifyResp{}
	user, err := jwt.ParseToken(req.GetToken())
	if err != nil {
		resp.Res = false
		resp.Code = http.StatusUnauthorized
		return &resp, nil
	}

	ok, err := casbin.Check(user.Subject, req.Method, req.Uri, true)
	if err != nil {
		resp.Res = false
		resp.Code = http.StatusInternalServerError
		return &resp, nil
	}

	if !ok {
		resp.Res = false
		resp.Code = http.StatusForbidden
		return &resp, nil
	}

	resp.Res = true
	resp.Code = 0
	resp.UserId = user.UserID
	return &resp, nil
}
