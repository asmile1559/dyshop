package main

import (
	service "github.com/asmile1559/dyshop/app/payment/biz/service"
	pbpayment "github.com/asmile1559/dyshop/pb/backend/payment"
	"golang.org/x/net/context"
)

type PaymentServiceServer struct {
	pbpayment.UnimplementedPaymentServiceServer
}

func (s *PaymentServiceServer) Charge(ctx context.Context, req *pbpayment.ChargeReq) (*pbpayment.ChargeResp, error) {
	resp, err := service.NewChargeService(ctx).Run(req)

	return resp, err
}
