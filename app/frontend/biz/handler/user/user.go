package user

import (
	"fmt"
	"os"

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
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		logrus.Error(errs.Translate(trans))
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": errs.Translate(trans),
		})
		return
	}

	// 业务逻辑
	req = user_page.RegisterReq{
		Email:           p.Email,
		Password:        p.Password,
		ConfirmPassword: p.ConfirmPassword,
	}
	_, err = service.NewRegisterService(c).Run(&req)

	if err != nil {
		logrus.WithError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "register ok!",
	})
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
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		logrus.Error(errs.Translate(trans))
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": errs.Translate(trans),
		})
		return
	}

	// 业务逻辑
	req = user_page.LoginReq{
		Email:    p.Email,
		Password: p.Password,
	}

	resp, err := service.NewLoginService(c).Run(&req)

	if err != nil {
		logrus.WithError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "login ok!",
		"token":   resp["token"],
	})

}

func UpdateUserInfo(c *gin.Context) {
	var err error
	var req user_page.UpdateUserInfoReq

	// 获取userid
	userId, ok := c.Get("user_id")
	if !ok {
		logrus.Error("no user id error")
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "no user id error",
		})
		return
	}

	// 业务逻辑
	// 绑定JSON请求体
	if err := c.Bind(&req); err != nil {
		logrus.WithError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "invalid request body",
			"err":     err.Error(),
		})
		return
	}
	req.UserId = userId.(int64)

	// 调用 Service 层的业务逻辑
	resp, err := service.NewUpdateUserInfoService(c).Run(&req)

	// 错误处理
	if err != nil {
		logrus.WithError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "internal server error",
			"err":     err.Error(),
		})
		return
	}

	// 返回成功的响应
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "update info ok",
		"resp":    resp,
	})
}

func GetUserInfo(c *gin.Context) {
	var err error
	var req user_page.GetUserInfoReq

	// 获取userid
	userId, ok := c.Get("user_id")
	if !ok {
		logrus.Error("no user id error")
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "no user id error",
		})
		return
	}

	//i, err := strconv.Atoi(id)
	req.UserId = userId.(int64)
	resp, err := service.NewGetUserInfoService(c).Run(&req)

	if err != nil {
		logrus.WithError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "internal server error",
		})
		return
	}

	c.HTML(http.StatusOK, "info.html", gin.H{
		"PageRouter": PageRouter,
		"UserInfo":   resp,
	})
}

func GetAccountInfo(c *gin.Context) {
	var err error
	var req user_page.GetAccountInfoReq

	// 获取userid
	userId, ok := c.Get("user_id")
	if !ok {
		logrus.Error("no user id error")
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "no user id error",
		})
		return
	}

	req.UserId = userId.(int64)
	resp, err := service.NewGetAccountInfoService(c).Run(&req)

	if err != nil {
		logrus.WithError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "internal server error",
		})
		return
	}

	c.HTML(http.StatusOK, "info.html", gin.H{
		"PageRouter": PageRouter,
		"UserInfo":   resp,
	})
}

func UpdateAccount(c *gin.Context) {
	var err error
	var req user_page.UpdateAccountReq

	// 参数校验
	// 未完成
	userId, ok := c.Get("user_id")
	if !ok {
		logrus.Error("no user id error")
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "no user id error",
		})
		return
	}

	// 绑定JSON请求体
	if err := c.Bind(&req); err != nil {
		logrus.WithError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "invalid request body",
			"err":     err.Error(),
		})
		return
	}

	req.UserId = userId.(int64)
	resp, err := service.NewUpdateAccountService(c).Run(&req)

	if err != nil {
		logrus.WithError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "update info ok",
		"resp":    resp,
	})

}

func RegisterMerchant(c *gin.Context) {
	var err error
	var req user_page.RegisterMerchantReq
	// 获取userid
	userId, ok := c.Get("user_id")
	if !ok {
		logrus.Error("no user id error")
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "no user id error",
		})
		return
	}

	// 业务逻辑
	req = user_page.RegisterMerchantReq{
		UserId: userId.(int64),
	}
	_, err = service.NewRegisterMerchantService(c).Run(&req)

	if err != nil {
		logrus.WithError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "update info ok",
		"resp": gin.H{
			"user_id": userId,
		},
	})

}

func UploadAvatar(c *gin.Context) {
	var req user_page.UploadAvatarReq

	if userId, ok := c.Get("user_id"); ok {
		file, err := c.FormFile("Img")
		if err != nil {
			logrus.WithError(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "upload file error.",
				"error":   err.Error(),
			})
			return
		}

		fileDir := fmt.Sprintf("/static/src/user/%v/", userId)
		saveDir := "." + fileDir
		if _, err := os.Stat(saveDir); os.IsNotExist(err) {
			err = os.Mkdir(saveDir, 0755)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"message": "something went wrong, please try it later.",
					"error":   err.Error(),
				})
				return
			}
		}

		//filePath := fileDir + file.Filename
		savePath := saveDir + file.Filename
		if err = c.SaveUploadedFile(file, savePath); err != nil {
			logrus.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "something went wrong, please try it later.",
				"error":   err.Error(),
			})
			return
		}

		req = user_page.UploadAvatarReq{
			UserId: userId.(int64),
			Url:    savePath,
		}
		resp, err := service.NewUploadAvatarService(c).Run(&req)

		if err != nil {
			logrus.WithError(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "internal server error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "upload ok!",
			"resp":    resp,
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"code":    http.StatusUnauthorized,
		"message": "no user id error",
	})
}

func Delete(c *gin.Context) {
	var err error
	var req user_page.DeleteUserReq

	// 获取userid
	userId, ok := c.Get("user_id")
	if !ok {
		logrus.Error("no user id error")
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "no user id error",
		})
		return
	}

	// 业务逻辑
	req = user_page.DeleteUserReq{
		UserId: userId.(int64),
	}
	_, err = service.NewDeleteUserService(c).Run(&req)

	if err != nil {
		logrus.WithError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "del account ok",
		"resp": gin.H{
			"user_id": userId,
		},
	})
}
