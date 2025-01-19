package user

import (
	u "github.com/dyshop/app/frontend/biz/handler/user"
	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine) {
	root := e.Group("/", _rootMw()...)
	{
		_user := root.Group("/user", _userMw()...)
		_user.POST("/login", append(_loginMw(), u.Login)...)
		_user.POST("/logout", append(_logoutMw(), u.Logout)...)
		_user.POST("/register", append(_registerMw(), u.Register)...)
	}
}
