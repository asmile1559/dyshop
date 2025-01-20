package checkout

import (
	"github.com/asmile1559/dyshop/app/frontend/middleware"
	"github.com/gin-gonic/gin"
)

func _rootMw() []gin.HandlerFunc {
	return nil
}

func _checkoutMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{middleware.Auth()}
}

func _checkout1Mw() []gin.HandlerFunc {
	return nil
}
