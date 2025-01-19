package user

import (
	"github.com/dyshop/app/frontend/biz/service"
	"github.com/dyshop/pb/frontend/user_page"

	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var err error
	var req user_page.LoginReq

	err = c.Bind(&req)
	if err != nil {
		c.String(http.StatusOK, "%v", err)
		return
	}

	_, err = service.NewLoginService(c).Run(&req)
	if err != nil {
		c.String(http.StatusOK, "%v", err)
		return
	}

	c.Redirect(http.StatusFound, "/ping")
}

func Logout(c *gin.Context) {

}

func Register(c *gin.Context) {

}
