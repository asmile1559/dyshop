package home

import (
	"fmt"
	"github.com/asmile1559/dyshop/app/frontend/biz/service"
	"github.com/asmile1559/dyshop/pb/frontend/home_page"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func Homepage(c *gin.Context) {
	req := home_page.GetHomepageReq{}

	ghpResp := service.NewGetHomepageService(c).Run(&req)

	resp := gin.H{
		"PageRouter":   pageRouter,
		"CategoryList": categoryList,
		"Carousels":    carousels,
		"Products":     ghpResp["resp"],
	}
	c.HTML(http.StatusOK, "index.html", &resp)
}

func GetShowcase(c *gin.Context) {
	which := c.Query("sub")
	if which == "" {
		which = "hot"
	}
	req := home_page.GetShowcaseReq{Which: which}
	gsResp := service.NewGetShowcaseService(c).Run(&req)

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"message":  fmt.Sprintf("获取新橱窗成功, 新的橱窗类别为: %v", which),
		"products": gsResp["resp"],
	})
}

func VerifyToken(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Authorization header is required",
		})
		return
	}

	token := strings.Split(authHeader, " ")
	if token[0] != "Bearer" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid Authorization header format. Expected 'Bearer <token>'",
		})
		return
	}
	vtResp := service.NewVerifyHomepageStatus(c).Run(&home_page.VerifyHomepageStatusReq{Token: token[1]})

	if vtResp == nil {
		logrus.Error("Verify Homepage Status Failed....")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "verify ok!",
		"resp":    vtResp["resp"],
	})
}
