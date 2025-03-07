package service

import (
	"github.com/asmile1559/dyshop/app/auth/utils/casbin"
	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
	"golang.org/x/net/context"
	"testing"
)

func init() {
	_ = casbin.InitEnforcer("./conf/model.conf", "./conf/policy.csv")
}

func TestVerifyRun(t *testing.T) {
	resp, _ := NewVerifyService(context.Background()).Run(&pbauth.VerifyTokenReq{
		Token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjUyMTMxNzgyNzUzMDM0MjQsIlN1YmplY3QiOiIxMjg1NWJhNzQwMTAwMCIsImlzcyI6ImR5c2hvcC1hdXRoIiwiZXhwIjoxNzQwNzc3ODI0LCJuYmYiOjE3NDA3MzQ2MjQsImlhdCI6MTc0MDczNDYyNH0.-yCJ5qtoLGPf9j_c7t1pJoqnaucNv0iZwfK8XIJUqG8",
		Method: "GET",
		Uri:    "/test/access",
	})

	if !resp.GetRes() {
		t.Fail()
		return
	}
	t.Log(resp)
}

// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlN1YmplY3QiOiIwMDAwMDAwMSIsImlzcyI6ImR5c2hvcC1nYXRld2F5IiwiZXhwIjoxNzY5MjIyMTI3LCJuYmYiOjE3Mzc2ODYxMjcsImlhdCI6MTczNzY4NjEyN30.w_yD3_YPNyAf71P1jn1LhMoiNv9Tp_B9ngjnHMyL3yw
