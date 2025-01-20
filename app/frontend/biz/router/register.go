package router

import (
	routercart "github.com/asmile1559/dyshop/app/frontend/biz/router/cart"
	routercheckout "github.com/asmile1559/dyshop/app/frontend/biz/router/checkout"
	routerorder "github.com/asmile1559/dyshop/app/frontend/biz/router/order"
	routerpayment "github.com/asmile1559/dyshop/app/frontend/biz/router/payment"
	routerproduct "github.com/asmile1559/dyshop/app/frontend/biz/router/product"
	routeruser "github.com/asmile1559/dyshop/app/frontend/biz/router/user"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(e *gin.Engine) {
	routeruser.Register(e)

	routerproduct.Register(e)

	routerorder.Register(e)

	routercart.Register(e)

	routerpayment.Register(e)

	routercheckout.Register(e)
}
