package service

import (
	"context"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbpayment "github.com/asmile1559/dyshop/pb/backend/payment"
	"github.com/asmile1559/dyshop/pb/frontend/payment_page"
	"github.com/gin-gonic/gin"
)

type ChargeService struct {
	ctx context.Context
}

func NewChargeService(c context.Context) *ChargeService {
	return &ChargeService{ctx: c}
}

func (s *ChargeService) Run(req *payment_page.ChargeReq) (map[string]interface{}, error) {

	// id, ok := s.ctx.Value("user_id").(uint32)
	// if !ok {
	// 	return nil, errors.New("expect user id")
	// }
	reqCred := req.GetCreditCard()

	paymentClient, conn, err := rpcclient.GetPaymentClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := paymentClient.Charge(s.ctx, &pbpayment.ChargeReq{
		TransactionId: req.TransactionId,
		CreditCard: &pbpayment.CreditCardInfo{
			CreditCardNumber:          reqCred.GetCreditCardNumber(),
			CreditCardCvv:             reqCred.GetCreditCardCvv(),
			CreditCardExpirationYear:  reqCred.GetCreditCardExpirationYear(),
			CreditCardExpirationMonth: reqCred.GetCreditCardExpirationMonth(),
		},
		FinalPrice: req.FinalPrice,
	})
	if err != nil {
		return nil, err
	}

	return gin.H{
		"transaction_id": resp.TransactionId,
	}, nil

	//return gin.H{
	//	"status": "charge ok",
	//}, nil
}
