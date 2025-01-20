package user

import (
	u "github.com/asmile1559/dyshop/app/frontend/biz/handler/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(e *gin.Engine) {
	root := e.Group("/", _rootMw()...)
	{
		_user := root.Group("/user", _userMw()...)
		_user.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", gin.H{})
		})
		_user.GET("/register", func(c *gin.Context) {
			c.HTML(http.StatusOK, "register.html", gin.H{})
		})
		//_user.GET("/logout", func(c *gin.Context) {
		// 	// clear token info in server
		//	c.Redirect(http.StatusFound, "/")
		//})
		_user.POST("/login", append(_loginMw(), u.Login)...)
		_user.POST("/register", append(_registerMw(), u.Register)...)
		//_user.POST("/logout", append(_logoutMw(), u.Logout)...)
	}
}
