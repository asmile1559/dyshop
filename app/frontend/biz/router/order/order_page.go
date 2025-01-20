package order

import (
	o "github.com/asmile1559/dyshop/app/frontend/biz/handler/order"
	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine) {
	root := e.Group("/", _rootMw()...)
	{
		_order := root.Group("/order", _orderMw()...)
		// TODO: order front page
		//_order.GET("/", )
		_order.POST("/place", append(_placeOrderMw(), o.PlaceOrder)...)
	}
}
