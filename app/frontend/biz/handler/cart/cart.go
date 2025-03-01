package cart

import (
	"net/http"

	"github.com/asmile1559/dyshop/app/frontend/biz/service"
	"github.com/asmile1559/dyshop/pb/frontend/cart_page"
	"github.com/gin-gonic/gin"
)

// AddItem 添加商品到购物车
func AddItem(c *gin.Context) {
	// GET from frontend
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "no user id error",
		})
		return
	}
	req := struct {
		ProductId   int    `json:"product_id"`
		ProductSpec gin.H  `json:"product_spec"`
		Quantity    int    `json:"quantity"`
		Postage     string `json:"postage"`
		Currency    string `json:"currency"`
	}{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "StatusBadRequest",
			"err":     err.Error(),
		})
		return
	}

	// reqJson to reqGrpc
	reqGrpc := cart_page.AddItemReq{
		ProductId: uint32(req.ProductId),
		Quantity:  int32(req.Quantity),
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := service.NewAddItemService(c).Run(&reqGrpc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// resp 是 gin.H{"resp": ...} 的结构，你也可以直接返回 resp["resp"]
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Add item success",
		"resp": gin.H{
			"user_id": userId,
			"content": resp,
		},
	})
}

// EmptyCart 清空购物车
func EmptyCart(c *gin.Context) {
	var req cart_page.EmptyCartReq
	// 如有需要从请求体获取参数, 可用 c.BindJSON(&req)
	// 目前该 Proto/Req 可能是空的

	resp, err := service.NewEmptyCartService(c).Run(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Empty cart success",
		"data":    resp,
	})
}

// GetCart 获取购物车
func GetCart(c *gin.Context) {
	var req cart_page.GetCartReq

	// 如果你需要从URL/Query里拿参数，可做: c.BindQuery(&req) 或 c.Bind(&req)
	// 这里暂时不需要, 直接调用后端

	resp, err := service.NewGetCartService(c).Run(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get cart success",
		"data":    resp, // resp 里通常是 { "resp": {Cart: {...}} }
	})
}

// DeleteCart 删除购物车
func DeleteCart(c *gin.Context) {
	var req cart_page.DeleteCartReq
	// 如有需要从请求体获取参数, 可用 c.BindJSON(&req)
	// 目前该 Proto/Req 可能是空的

	resp, err := service.NewDeleteCartService(c).Run(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete cart success",
		"data":    resp,
	})
}
