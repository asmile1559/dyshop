package payment

import (
	"github.com/asmile1559/dyshop/app/frontend/middleware"
	"github.com/gin-gonic/gin"
)

func _rootMw() []gin.HandlerFunc {
	return nil
}

func _paymentMW() []gin.HandlerFunc {
	return []gin.HandlerFunc{middleware.Auth()}
}

func _chargeMw() []gin.HandlerFunc {
	return nil
}
