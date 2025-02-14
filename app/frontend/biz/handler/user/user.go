package user

import (
	"strconv"

	"github.com/asmile1559/dyshop/app/frontend/biz/model"
	"github.com/asmile1559/dyshop/app/frontend/biz/service"
	"github.com/asmile1559/dyshop/pb/frontend/user_page"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var err error
	var req user_page.RegisterReq
	var p model.ParamRegister
	
	// 参数校验
	if err := c.Bind(&p); err != nil {
		//请求参数有误，返回响应
		logrus.WithError(err).Error("register with invalid param")
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			logrus.WithError(err)
			c.String(http.StatusOK, "An error occurred: %v", err)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		logrus.Error(errs.Translate(trans))
		c.String(http.StatusOK, "An error occurred: %v", errs.Translate(trans))
		return
	}

	// 业务逻辑
	req = user_page.RegisterReq{
		Email:           p.Email,
		Password:        p.Password,
		ConfirmPassword: p.ConfirmPassword,
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
	var p model.ParamLogin
	
	// 参数校验
	if err := c.Bind(&p); err != nil {
		//请求参数有误，返回响应
		logrus.WithError(err).Error("register with invalid param")
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			logrus.WithError(err)
			c.String(http.StatusOK, "An error occurred: %v", err)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		logrus.Error(errs.Translate(trans))
		c.String(http.StatusOK, "An error occurred: %v", errs.Translate(trans))
		return
	}

	// 业务逻辑
	req = user_page.LoginReq{
		Email:    p.Email,
		Password: p.Password,
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
	var p model.ParamUpdateUser
	
	// 参数校验
	if err := c.Bind(&p); err != nil {
		//请求参数有误，返回响应
		logrus.WithError(err).Error("register with invalid param")
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			logrus.WithError(err)
			c.String(http.StatusOK, "An error occurred: %v", err)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		logrus.Error(errs.Translate(trans))
		c.String(http.StatusOK, "An error occurred: %v", errs.Translate(trans))
		return
	}

	// 业务逻辑
	req = user_page.UpdateUserReq{
		UserId:   p.UserID,
		Email:    p.Email,
		Password: p.Password,
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
	req.UserId = int64(i)
	
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
	var p model.ParamDeleteUser
	
	// 参数校验
	if err := c.Bind(&p); err != nil {
		//请求参数有误，返回响应
		logrus.WithError(err).Error("register with invalid param")
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			logrus.WithError(err)
			c.String(http.StatusOK, "An error occurred: %v", err)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		logrus.Error(errs.Translate(trans))
		c.String(http.StatusOK, "An error occurred: %v", errs.Translate(trans))
		return
	}

	// 业务逻辑
	req = user_page.DeleteUserReq{
		UserId: p.UserID,
	}
	resp, err := service.NewDeleteUserService(c).Run(&req)

	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	c.String(http.StatusOK, "%v", resp)
}