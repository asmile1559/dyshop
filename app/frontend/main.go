package main

import (
	bizrouter "github.com/dyshop/app/frontend/biz/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterInfo struct {
	Email      string `form:"email" binding:"required"`
	Password   string `form:"password" binding:"required"`
	RePassword string `form:"re_password" binding:"required,eqfield=Password"`
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/**")
	router.StaticFS("/static", http.Dir("static"))

	router.GET("/ping", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pong.html", gin.H{
			"host": "192.168.191.130:10166",
			"ping": "pong",
		})
	})

	bizrouter.RegisterRouters(router)
	//router.GET("/register", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "register.html", gin.H{})
	//})
	//
	//router.POST("/register", func(c *gin.Context) {
	//	registerInfo := RegisterInfo{}
	//	err := c.Bind(&registerInfo)
	//	if err != nil {
	//		c.JSON(http.StatusOK, gin.H{
	//			"err": err,
	//		})
	//		return
	//	}
	//	c.String(http.StatusOK, "congratulation! you have registered! your email is %v", registerInfo.Email)
	//})
	router.Run(":10166")
}
