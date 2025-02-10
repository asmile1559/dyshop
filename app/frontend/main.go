package main

import (
	"net/http"

	"github.com/asmile1559/dyshop/app/frontend/middleware"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/asmile1559/dyshop/utils/jwt"
	"github.com/sirupsen/logrus"

	bizrouter "github.com/asmile1559/dyshop/app/frontend/biz/router"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	hookx.Init(hookx.DefaultHook)
}

func main() {
	rpcclient.InitRPCClient()

	router := gin.Default()

	router.LoadHTMLGlob("templates/**")
	router.StaticFS("/static", http.Dir("static"))

	router.GET("/ping", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pong.html", gin.H{
			"code": http.StatusOK,
			"host": "192.168.191.130:10166",
			"ping": "pong",
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "404 page not found",
		})
	})

	bizrouter.RegisterRouters(router)
	registerTestRouter(router)

	if err := router.Run(":" + viper.GetString("server.port")); err != nil {
		logrus.Fatal(err)
	}
	//router.Run(":10166")
}

func registerTestRouter(e *gin.Engine) {
	type user struct {
		UserID uint32 `json:"user_id"`
	}

	mid := func(c *gin.Context) {
		logrus.Infof("Method: %v, URI: %v", c.Request.Method, c.Request.RequestURI)
		c.Next()
	}

	_test := e.Group("/test")
	_test.POST("/login", func(c *gin.Context) {
		u := user{}
		err := c.BindJSON(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Expect json format login information",
			})
			return
		}
		token, err := jwt.GenerateJWT(u.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Something went wrong. Please try again later.",
			})
			logrus.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "login ok!",
			"token":   token,
		})

	})
	_test.GET("/access", mid, middleware.Auth(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Auth Test OK",
		})
	})
	_test.POST("/access", mid, middleware.Auth(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Auth Test OK",
		})
	})
}
