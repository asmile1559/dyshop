package cart

import (
	"github.com/asmile1559/dyshop/app/frontend/biz/service"
	"github.com/asmile1559/dyshop/pb/frontend/cart_page"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddItem(c *gin.Context) {
	var req cart_page.AddItemReq

	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	resp, err := service.NewAddItemService(c).Run(&req)

	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	c.String(http.StatusOK, "GetProduct ok! your id is: %v", resp)
}

func EmptyCart(c *gin.Context) {
	var req cart_page.EmptyCartReq

	resp, err := service.NewEmptyCartService(c).Run(&req)

	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	// redirect to /cart
	c.String(http.StatusOK, "GetProduct ok! your id is: %v", resp)
}

func GetCart(c *gin.Context) {
	panic("DO NOT use the function! Use DeliverTokenService directly")
}

//func ListCart(c *gin.Context) {
//	// ListCart API is not provided
//	panic("DO NOT use the function! Use DeliverTokenService directly")
//}
