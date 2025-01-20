package order

import (
	"github.com/asmile1559/dyshop/app/frontend/middleware"
	"github.com/gin-gonic/gin"
)

func _rootMw() []gin.HandlerFunc {
	return nil
}

func _orderMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{middleware.Auth()}
}

func _listOrdersMw() []gin.HandlerFunc {
	return nil
}

func _placeOrderMw() []gin.HandlerFunc {
	return nil
}

func _markOrderPaidMw() []gin.HandlerFunc {
	return nil
}

func _getOrderMw() []gin.HandlerFunc {
	return nil
}

func _modifyOrder() []gin.HandlerFunc {
	return nil
}
func _cancelOrderMw() []gin.HandlerFunc {
	return nil
}
