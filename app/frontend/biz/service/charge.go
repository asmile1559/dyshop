package service

import (
	"context"
	"github.com/asmile1559/dyshop/pb/frontend/payment_page"
	"github.com/gin-gonic/gin"
)

//	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
//	pbpayment "github.com/asmile1559/dyshop/pb/backend/payment"

type ChargeService struct {
	ctx context.Context
}

func NewChargeService(c context.Context) *ChargeService {
	return &ChargeService{ctx: c}
}

func (s *ChargeService) Run(req *payment_page.ChargeReq) (map[string]interface{}, error) {

	//id, ok := s.ctx.Value("user_id").(uint32)
	//if !ok {
	//	return nil, errors.New("expect user id")
	//}
	//reqCred := req.GetCreditCard()
	//
	//resp, err := rpcclient.PaymentClient.Charge(s.ctx, &pbpayment.ChargeReq{
	//	Amount: req.GetAmount(),
	//	CreditCard: &pbpayment.CreditCardInfo{
	//		CreditCardNumber:          reqCred.GetCreditCardNumber(),
	//		CreditCardCvv:             reqCred.GetCreditCardCvv(),
	//		CreditCardExpirationYear:  reqCred.GetCreditCardExpirationYear(),
	//		CreditCardExpirationMonth: reqCred.GetCreditCardExpirationMonth(),
	//	},
	//	OrderId: req.GetOrderId(),
	//	UserId:  id,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//return gin.H{
	//	"resp": resp,
	//}, nil

	return gin.H{
		"status": "charge ok",
	}, nil
}
