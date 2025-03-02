package user

import (
	"io"

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
		logrus.WithError(err).Error("register error")
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

	c.HTML(http.StatusOK, "account.html", gin.H{
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
		fileHeader, err := c.FormFile("Img")
		if err != nil {
			logrus.WithError(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "upload file error.",
				"error":   err.Error(),
			})
			return
		}

		// **文件大小检查（如果超过 4MB，返回错误）**
		const maxFileSize = 4 * 1024 * 1024 // 4MB
		if fileHeader.Size > maxFileSize {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "File size exceeds 4MB limit",
				"error":   "File size exceeds 4MB limit",
			})
			return
		}
		// 打开文件
		file, err := fileHeader.Open()
		if err != nil {
			logrus.WithError(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "upload file error.",
				"error":   err.Error(),
			})
			return
		}
		defer file.Close()

		// 读取整个文件
		imageData, err := io.ReadAll(file)
		if err != nil {
			logrus.WithError(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "read file error.",
				"error":   err.Error(),
			})
			return
		}

		// 发送请求
		req = user_page.UploadAvatarReq{
			UserId:    userId.(int64),
			Filename:  fileHeader.Filename,
			ImageData: imageData, // **一次性发送完整数据**
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
