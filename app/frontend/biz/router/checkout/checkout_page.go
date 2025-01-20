package checkout

import (
	c "github.com/asmile1559/dyshop/app/frontend/biz/handler/checkout"
	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine) {
	_root := e.Group("/", _rootMw()...)
	{
		_checkout := _root.Group("/checkout", _checkoutMw()...)
		// TODO: checkout front page
		//_checkout.GET("/")
		_checkout.POST("/checkout", append(_checkout1Mw(), c.Checkout)...)
	}
}
