package user

import (
	"strconv"

	"github.com/asmile1559/dyshop/app/frontend/biz/service"
	"github.com/asmile1559/dyshop/pb/frontend/user_page"

	"net/http"

	"github.com/gin-gonic/gin"
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

}

func UpdateUser(c *gin.Context) {
	var err error
	var req user_page.UpdateUserReq

	// 绑定请求数据
	err = c.Bind(&req)
	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	// 调用 Service 层的业务逻辑
	resp, err := service.NewUpdateUserService(c).Run(&req)

	// 错误处理
	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	// 返回成功的响应
	c.String(http.StatusOK, "%v", resp)
}

func GetUserInfo(c *gin.Context){
	var err error
	var req user_page.GetUserInfoReq

	id := c.Param("id")
	if id == "" {
		c.String(http.StatusOK, "expect a product index")
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}
	req.UserId = uint32(i)
	
	resp, err := service.NewGetUserInfoService(c).Run(&req)

	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	c.String(http.StatusOK, "%v", resp)
}

func Delete(c *gin.Context){
	var err error
	var req user_page.DeleteUserReq

	err = c.Bind(&req)
	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	resp, err := service.NewDeleteUserService(c).Run(&req)

	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	c.String(http.StatusOK, "%v", resp)
}