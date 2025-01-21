package service

import (
	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
	"golang.org/x/net/context"
	"testing"
)

func TestVerifyRun(t *testing.T) {
	resp, err := NewVerifyService(context.Background()).Run(&pbauth.VerifyTokenReq{Token: "123"})
	if err != nil {
		t.Fail()
		return
	}
	if !resp.GetRes() {
		t.Fail()
		return
	}
	t.Log(resp)
}
