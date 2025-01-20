package user_router

import (
	"github.com/asmile1559/dyshop/app/frontend/internal/controller/user_controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoute(e *gin.Engine) {
	root := e.Group("/", _rootMw()...)
	{
		_user := root.Group("/user", _userMw()...)
		_user.GET("/register", func(c *gin.Context) {
			c.HTML(http.StatusOK, "register.html", nil)
		})
		_user.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", nil)
		})
		_user.POST("/register", append(_registerMw(), user_controller.Register)...)
		_user.POST("/login", append(_loginMw(), user_controller.Login)...)
		//_user.POST("/logout", append(_logoutMw(), user_controller.Logout)...)
	}
}
