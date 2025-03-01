package product

import (
	"encoding/json"
	"fmt"
	"github.com/asmile1559/dyshop/app/frontend/biz/service"
	"github.com/asmile1559/dyshop/pb/frontend/product_page"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"os"
	"strconv"
)

func GetProduct(c *gin.Context) {
	var err error
	var req product_page.GetProductReq
	userId, _ := c.Get("user_id")
	productId := c.Query("product_id")

	if productId == "" {
		c.String(http.StatusOK, "expect a product index")
		return
	}

	i, err := strconv.Atoi(productId)
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

	resp1 := gin.H{
		"PageRouter": PageRouter,
		"UserInfo": gin.H{
			"Id":   userId,
			"Name": "lixiaoming",
			"Sign": "许仙给老婆买了一顶帽子，白娘子戴上之后就死了，因为那是顶鸭（压）舌（蛇）帽。",
			"Img":  "/static/src/user/snake.svg",
			"Role": []string{"user", "merchant"},
		},
		"Products": []gin.H{resp},
	}
	c.HTML(http.StatusOK, "product-page.html", resp1)
}

func SearchProduct(c *gin.Context) {
	var err error
	var req product_page.SearchProductsReq
	kw := c.Query("keyword")
	category := c.Query("category")
	if kw == "" && category == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "must have a keyword or a category",
		})
		return
	}
	pg := c.Query("pg")
	if pg == "" {
		pg = "1"
	}
	sort := c.Query("sort")
	if sort == "" {
		sort = "comprehensive"
	}
	ps := c.Query("ps")
	if ps == "" {
		ps = "30"
	}
	curPage, err := strconv.Atoi(pg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "page must be a number",
			"error":   err.Error(),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "page must be a number",
			"error":   err.Error(),
		})
		return
	}
	if curPage <= 0 {
		curPage = 1
	}
	pagesize, err := strconv.Atoi(ps)
	//totalPage := 8
	req.Page = int32(curPage)
	req.Query = kw
	req.Category = category
	req.Sort = sort
	req.PageSize = int32(pagesize)
	resp, err := service.NewSearchProductService(c).Run(&req)

	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	c.HTML(http.StatusOK, "search.html", resp)
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
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "no user id error",
		})
		return
	}

	req1 := struct {
		ProductId         string           `json:"product_id"`
		ProductImg        string           `json:"product_img"`
		ProductName       string           `json:"product_name"`
		ProductDesc       string           `json:"product_desc"`
		ProductSold       string           `json:"product_sold"`
		ProductSpecs      []map[string]any `json:"product_specs"`
		ProductCategories []string         `json:"product_categories"`
		ProductParams     []map[string]any `json:"product_params"`
		ProductInsurance  string           `json:"product_insurance"`
		ProductExpress    string           `json:"product_express"`
	}{}

	err := json.Unmarshal([]byte(c.PostForm("product")), &req1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "upload file error. 1",
			"error":   err.Error(),
		})
		return
	}
	if req1.ProductImg == "" {
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "upload file error. 2",
				"error":   err.Error(),
			})
			return
		}

		fileDir := fmt.Sprintf("/static/src/product/%v/", req1.ProductId)
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
		filePath := fileDir + file.Filename
		savePath := saveDir + file.Filename
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			logrus.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "something went wrong, please try it later.",
				"error":   err.Error(),
			})
			return
		}
		req1.ProductImg = filePath
	}

	if req1.ProductId == "0" {
		req1.ProductId = "3"
	}
	// 绑定JSON请求体
	if err := c.ShouldBindJSON(&req1); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	req.Name = &req1.ProductName
	req.Description = &req1.ProductDesc
	req.Categories = req1.ProductCategories
	req.Picture = &req1.ProductImg
	req.Price = nil
	// 调用服务层
	_, err = service.NewModifyProductService(c.Request.Context()).Run(&req)
	if err != nil {
		handleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "upload ok!",
		"resp": gin.H{
			"user_id": userId,
			"product": req,
		},
	})
	return

}

// DeleteProduct 删除产品 Handler
func DeleteProduct(c *gin.Context) {
	// 从查询参数获取ID
	idStr := c.Query("product_id")
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
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "no user id error",
		})
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

	fmt.Println(req)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "del account ok",
		"resp": gin.H{
			"user_id":    userId,
			"product_id": &req.Id,
		},
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
