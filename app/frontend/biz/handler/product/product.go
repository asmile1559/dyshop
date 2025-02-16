package product

import (
	"github.com/asmile1559/dyshop/app/frontend/biz/service"
	"github.com/asmile1559/dyshop/pb/frontend/product_page"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

// CreateProduct 创建产品 Handler
func CreateProduct(c *gin.Context) {
	var req product_page.CreateProductReq

	// 绑定JSON请求体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// 调用服务层
	resp, err := service.NewCreateProductService(c.Request.Context()).Run(&req)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    0,
		"message": "created successfully",
		"data":    resp,
	})
}

// ModifyProduct 修改产品 Handler
func ModifyProduct(c *gin.Context) {
	var req product_page.ModifyProductReq

	// 绑定JSON请求体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// 调用服务层
	resp, err := service.NewModifyProductService(c.Request.Context()).Run(&req)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "updated successfully",
		"data":    resp,
	})
}

// DeleteProduct 删除产品 Handler
func DeleteProduct(c *gin.Context) {
	// 从查询参数获取ID
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing id parameter"})
		return
	}

	// 转换ID类型
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	// 构造请求
	req := &product_page.DeleteProductReq{
		Id: uint32(id),
	}

	// 调用服务层
	_, err = service.NewDeleteProductService(c.Request.Context()).Run(req)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "deleted successfully",
	})
}

// 统一错误处理函数
func handleServiceError(c *gin.Context, err error) {
	if sts, ok := status.FromError(err); ok {
		switch sts.Code() {
		case codes.InvalidArgument:
			c.JSON(http.StatusBadRequest, gin.H{"error": sts.Message()})
		case codes.NotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": sts.Message()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
