package user

import (
	"github.com/asmile1559/dyshop/app/frontend/biz/service"
	"github.com/asmile1559/dyshop/pb/frontend/user_page"

	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var err error
	var req user_page.RegisterReq

	err = c.Bind(&req)
	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	resp, err := service.NewRegisterService(c).Run(&req)

	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	c.String(http.StatusOK, "%v", resp)
}

func Login(c *gin.Context) {
	var err error
	var req user_page.LoginReq

	err = c.Bind(&req)
	if err != nil {
		c.String(http.StatusOK, "%v", err)
		return
	}

	resp, err := service.NewLoginService(c).Run(&req)

	if err != nil {
		c.String(http.StatusOK, "%v", err)
		return
	}

	c.String(http.StatusOK, "%v", resp)

	//c.Redirect(http.StatusFound, "/")
}

// optional function
//func Logout(c *gin.Context) {
//
//}
