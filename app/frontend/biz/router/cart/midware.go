package cart

import (
	"github.com/asmile1559/dyshop/app/frontend/middleware"
	"github.com/gin-gonic/gin"
)

func _rootMw() []gin.HandlerFunc {
	return nil
}

func _cartMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{middleware.Auth()}
}

func _addCartMw() []gin.HandlerFunc {
	return nil
}

func _listCartMw() []gin.HandlerFunc {
	return nil
}

func _emptyCartMw() []gin.HandlerFunc {
	return nil
}

func _getCartMw() []gin.HandlerFunc {
	return nil
}
