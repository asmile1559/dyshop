package product

import (
	"github.com/asmile1559/dyshop/app/frontend/biz/service"
	"github.com/asmile1559/dyshop/pb/frontend/product_page"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetProduct(c *gin.Context) {
	var err error
	var req product_page.GetProductReq

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
	req.Id = uint32(i)

	resp, err := service.NewGetProductService(c).Run(&req)

	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	c.String(http.StatusOK, "GetProduct ok! your id is: %v", resp)
}

func SearchProduct(c *gin.Context) {
	var err error
	var req product_page.SearchProductsReq

	req.Q = c.Query("q")

	resp, err := service.NewSearchProductService(c).Run(&req)

	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	c.String(http.StatusOK, "SearchProduct ok! your id is: %v", resp)
}

func ListProduct(c *gin.Context) {
	panic("DO NOT use the function! Use ListProductService directly")
}
