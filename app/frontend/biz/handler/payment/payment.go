package payment

import (
	"github.com/asmile1559/dyshop/app/frontend/biz/service"
	"github.com/asmile1559/dyshop/pb/frontend/payment_page"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"strconv"
)

func Charge(c *gin.Context) {

	var err error
	//var req payment_page.ChargeReq
	req := struct {
		TransactionId string `json:"transaction_id"`
		CreditCard    gin.H  `json:"credit_card"`
		FinalPrice    string `json:"final_price"`
	}{}
	err = c.BindJSON(&req)
	fmt.Println("支付请求：", req)
	
	chargeReq := payment_page.ChargeReq{
		TransactionId: req.TransactionId,
		CreditCard:    &payment_page.CreditCardInfo{
			CreditCardNumber:          req.CreditCard["card_number"].(string),
			CreditCardCvv:             int32(mustStrToInt(req.CreditCard["cvv"].(string))),
			CreditCardExpirationYear:  int32(mustStrToInt(req.CreditCard["expire_year"].(string))),
			CreditCardExpirationMonth: int32(mustStrToInt(req.CreditCard["expire_month"].(string))),
		},
		FinalPrice:    req.FinalPrice,
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "StatusBadRequest",
			"err":     err.Error(),
		})
		return
	}

	resp, err := service.NewChargeService(c).Run(&chargeReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "StatusBadRequest",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "charge ok!",
		"resp": resp,
	})
}

func mustStrToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		// 处理转换错误，如果需要可以返回默认值或处理错误
		panic(err)
	}
	return val
}
