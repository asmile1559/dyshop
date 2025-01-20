package cart

import (
	c "github.com/asmile1559/dyshop/app/frontend/biz/handler/cart"
	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine) {
	root := e.Group("/", _rootMw()...)
	{
		_cart := root.Group("/cart", _cartMw()...)
		// TODO: cart front page
		//_cart.GET("/", )
		_cart.POST("/add", append(_addCartMw(), c.AddCart)...)
		_cart.GET("/empty", append(_emptyCartMw(), c.EmptyCart)...)
	}
}
