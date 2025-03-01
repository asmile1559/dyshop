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
		_cart.GET("/", append(_getCartMw(), c.GetCart)...)
		_cart.POST("/add", append(_addCartMw(), c.AddItem)...)
		_cart.POST("/empty", append(_emptyCartMw(), c.EmptyCart)...)
		_cart.POST("/delete", append(_deleteCartMw(), c.DeleteCart)...)
		// TODO: cart checkout page
	}
}
