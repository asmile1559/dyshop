package user_controller

import (
	"net/http"
	"user/internal/model/user_model/dto"
	"user/internal/service"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req dto.RegisterRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	_, err = service.NewUserService(c).Register(&req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "register ok!"})
}

func Login(c *gin.Context) {
	var req dto.LoginRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	resp, err := service.NewUserService(c).Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": resp})
}

//func Logout(c *gin.Context) {
//
//}
