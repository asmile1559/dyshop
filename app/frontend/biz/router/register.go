package router

import (
	routeruser "github.com/dyshop/app/frontend/biz/router/user"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(e *gin.Engine) {
	routeruser.Register(e)
}
