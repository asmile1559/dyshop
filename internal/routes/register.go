package routes

import (
	"github.com/asmile1559/dyshop/internal/routes/user_router"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	user_router.RegisterUserRoute(e)

}
