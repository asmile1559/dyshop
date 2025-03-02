package main

import (
	"html/template"
	"net/http"
	"unicode/utf8"

	"github.com/asmile1559/dyshop/app/frontend/biz/handler/user"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/sirupsen/logrus"

	bizrouter "github.com/asmile1559/dyshop/app/frontend/biz/router"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	feutils "github.com/asmile1559/dyshop/app/frontend/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	hookx.Init(hookx.DefaultHook)
}

func main() {
	rpcclient.InitRPCClient()

	//初始化gin框架内置的校验器翻译器
	if err := user.InitTrans("zh"); err != nil {
		logrus.Fatal(err)
	}

	router := gin.Default()
	router.Use(cors.Default())
	router.SetFuncMap(template.FuncMap{
		"realLen":    utf8.RuneCountInString,
		"iMod":       feutils.IndexMod,
		"rangeSlice": feutils.RangeSlice,
		"showPages":  feutils.ShowPages,
		"subOne":     feutils.SubOne,
		"addOne":     feutils.AddOne,
		"inSlice":    feutils.InSlice[string],
		"rowIdx":     feutils.RowIdx,
		"colIdx":     feutils.ColIdx,
		"calcPrice":  feutils.CalcPrice,
	})
	router.LoadHTMLGlob("templates/**")

	router.StaticFS("/static", http.Dir("static"))
	router.StaticFS("/example/static", http.Dir("static"))

	router.GET("/ping", func(c *gin.Context) {
		// pong.html
		//  Hi, this is pong page.
		//  <h1>I am {{ .Host }}</h1>
		//  <h2>{{ .Pong }}...</h2>

		//  1. 方式 1
		//resp := struct {
		//	Code int    `json:"code"`
		//	Host string `json:"host"`
		//	Pong string `json:"pong"`
		//}{http.StatusOK, "192.168.191.130:10166", "Pong"}
		//c.HTML(http.StatusOK, "pong.html", &resp)
		// 2. 方式 2
		c.HTML(http.StatusOK, "pong.html", gin.H{
			"Code": http.StatusOK,
			"Host": "localhost:10166",
			"Pong": "Pong",
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "404 page not found",
		})
	})

	bizrouter.RegisterRouters(router)
	registerExampleRouter(router)

	if err := router.Run(":" + viper.GetString("server.port")); err != nil {
		logrus.Fatal(err)
	}
	//router.Run(":10166")
}
