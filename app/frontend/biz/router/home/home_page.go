package home

import (
	h "github.com/asmile1559/dyshop/app/frontend/biz/handler/home"
	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine) {
	_home := e.Group("/")
	{
		_home.GET("/", h.Homepage)

		_home.POST("/verify", h.VerifyToken)

		_home.GET("/showcase", h.GetShowcase)
	}
}
