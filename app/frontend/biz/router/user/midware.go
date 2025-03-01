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

func _registerMw() []gin.HandlerFunc {
	return nil
}

func _updateuserinfoMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{middleware.Auth()}
}

func _getuserinfoMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{middleware.Auth()}
}

func _updateaccountinfoMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{middleware.Auth()}
}

func _getaccountinfoMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{middleware.Auth()}
}

func _uploadavatarMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{middleware.Auth()}
}

func _registermerchantMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{middleware.Auth()}
}

func _deleteMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{middleware.Auth()}
}