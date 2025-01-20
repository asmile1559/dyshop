package payment

import (
	p "github.com/asmile1559/dyshop/app/frontend/biz/handler/payment"
	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine) {
	_root := e.Group("/", _rootMw()...)
	{
		_payment := _root.Group("/payment", _paymentMW()...)
		// TODO: payment front page
		// _payment.Get("/", )
		_payment.POST("/charge", append(_chargeMw(), p.Charge)...)
	}
}
