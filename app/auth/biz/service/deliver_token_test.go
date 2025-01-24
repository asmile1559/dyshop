package service

import (
	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
	"golang.org/x/net/context"
	"testing"
)

//	func (s *DeliverService) Run(req *pbauth.DeliverTokenReq) (*pbauth.DeliveryResp, error) {
//		// TODO: finish your business code...
//		//
//		return &pbauth.DeliveryResp{Token: "deliver token ok!"}, nil
//	}

func TestDeliverTokenRun(t *testing.T) {
	resp, err := NewDeliverService(context.Background()).Run(&pbauth.DeliverTokenReq{UserId: 1})
	if err != nil {
		t.Fail()
		return
	}
	t.Log(resp.Token)
}
