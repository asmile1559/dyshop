package main

import (
	"log"
	"user/internal/dao/user_dao"
	"user/internal/routes"
	"utils/logx"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	logx.Init()
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	if err := initDB(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("static/pages/**")

	routes.RegisterRoutes(router)

	//router.Run(":10166")
	port := viper.GetString("server.port")
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("start server failed: %v", err)
	}
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	return viper.ReadInConfig()
}

// need to be separate to microservices
func initDB() error {
	return user_dao.InitDB()
}
