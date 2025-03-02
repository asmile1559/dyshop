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

	resp, err := service.NewAddItemService(c).Run(&reqGrpc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "no user id error",
		})
		return
	}

	reqGrpc := cart_page.EmptyCartReq{}

	resp, err := service.NewEmptyCartService(c).Run(&reqGrpc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Empty cart success",
		"resp": gin.H{
			"user_id": userId,
			"content": resp,
		},
	})
}

// GetCart 获取购物车
func GetCart(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "no user id error",
		})
		return
	}

	reqGrpc := cart_page.GetCartReq{}

	resp, userinfo, err := service.NewGetCartService(c).Run(&reqGrpc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "cart.html", gin.H{
		"PageRouter": PageRouter,
		"UserInfo": gin.H{
			"UserId": userId,
			"Name":   userinfo["Name"],
		},
		"CartItems": resp,
	})
}

// DeleteCart 删除购物车部分条目
func DeleteCart(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "no user id error",
		})
		return
	}

	req := struct {
		ItemIds []uint32 `json:"item_ids"`
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

	reqGrpc := cart_page.DeleteCartReq{
		ItemIds: make([]uint32, 0, len(req.ItemIds)),
	}
	reqGrpc.ItemIds = append(reqGrpc.ItemIds, req.ItemIds...)

	resp, err := service.NewDeleteCartService(c).Run(&reqGrpc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete cart success",
		"resp": gin.H{
			"user_id": userId,
			"content": resp,
		},
	})
}
