package user

import (
	"github.com/asmile1559/dyshop/app/frontend/middleware"
	"github.com/gin-gonic/gin"
)

func _rootMw() []gin.HandlerFunc {
	return nil
}

func _userMw() []gin.HandlerFunc {
	return nil
}

func _loginMw() []gin.HandlerFunc {
	return nil
}

func _logoutMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{middleware.Auth()}
}

func _registerMw() []gin.HandlerFunc {
	return nil
}

func _updateMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{middleware.Auth()}
}

func _infoMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{middleware.Auth()}
}

func _deleteMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{middleware.Auth()}
}