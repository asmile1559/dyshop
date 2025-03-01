package cart

import (
	"github.com/gin-gonic/gin"
)

func _rootMw() []gin.HandlerFunc {
	return nil
}

func _cartMw() []gin.HandlerFunc {
	// return []gin.HandlerFunc{middleware.Auth()}
	return nil
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

func _deleteCartMw() []gin.HandlerFunc {
	return nil
}
