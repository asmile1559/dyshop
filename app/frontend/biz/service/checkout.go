package service

import (
	"context"
	"github.com/asmile1559/dyshop/pb/frontend/checkout_page"
	"github.com/gin-gonic/gin"
)

//rpccleint "github.com/asmile1559/dyshop/app/frontend/rpc"
//pbcheckout "github.com/asmile1559/dyshop/pb/backend/checkout"
//pbpayment "github.com/asmile1559/dyshop/pb/backend/payment"

type CheckoutService struct {
	ctx context.Context
}

func NewCheckoutService(c context.Context) *CheckoutService {
	return &CheckoutService{ctx: c}
}

func (s *CheckoutService) Run(req *checkout_page.CheckoutReq) (map[string]interface{}, error) {

	//var id uint32
	//var ok bool
	//
	//if id, ok = s.ctx.Value("user_id").(uint32); ok {
	//	return nil, errors.New("expect user id")
	//}
	//
	//reqAddr := req.GetAddress()
	//reqCred := req.GetCreditCard()
	//
	//resp, err := rpccleint.CheckoutClient.Checkout(s.ctx, &pbcheckout.CheckoutReq{
	//	UserId:    id,
	//	Firstname: req.GetFirstname(),
	//	Lastname:  req.GetLastname(),
	//	Email:     req.GetEmail(),
	//	Address: &pbcheckout.Address{
	//		StreetAddress: reqAddr.GetStreetAddress(),
	//		City:          reqAddr.GetCity(),
	//		State:         reqAddr.GetState(),
	//		Country:       reqAddr.GetCountry(),
	//		ZipCode:       reqAddr.GetZipCode(),
	//	},
	//	CreditCard: &pbpayment.CreditCardInfo{
	//		CreditCardNumber:          reqCred.GetCreditCardNumber(),
	//		CreditCardCvv:             reqCred.GetCreditCardCvv(),
	//		CreditCardExpirationYear:  reqCred.GetCreditCardExpirationYear(),
	//		CreditCardExpirationMonth: reqCred.GetCreditCardExpirationMonth(),
	//	},
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//return gin.H{
	//	"resp": resp,
	//}, nil

	return gin.H{
		"status": "checkout ok",
	}, nil
}
