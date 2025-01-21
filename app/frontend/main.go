package main

import (
	"github.com/sirupsen/logrus"
	"log"
	"net/http"

	bizrouter "github.com/asmile1559/dyshop/app/frontend/biz/router"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	"github.com/asmile1559/dyshop/utils/logx"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	initLog()

	rpcclient.InitRPCClient()

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

	if err := router.Run(":" + viper.GetString("server.port")); err != nil {
		logrus.Fatal(err)
	}
	//router.Run(":10166")
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	return viper.ReadInConfig()
}

func initLog() {
	logx.Init()
}
