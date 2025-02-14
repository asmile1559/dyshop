package user

import (
	"net/http"

	u "github.com/asmile1559/dyshop/app/frontend/biz/handler/user"
	"github.com/gin-gonic/gin"
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
		_user.POST("/login", append(_loginMw(), u.Login)...)
		_user.POST("/register", append(_registerMw(), u.Register)...)
		_user.PUT("/update", append(_updateMw(),u.UpdateUser)...)
		_user.GET("/info/:id", append(_infoMw(),u.GetUserInfo)...)
		_user.DELETE("/delete", append(_deleteMw(),u.Delete)...)
	}
}
