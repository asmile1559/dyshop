package order

import (
	o "github.com/asmile1559/dyshop/app/frontend/biz/service"
	"github.com/asmile1559/dyshop/pb/frontend/order_page"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOrders(c *gin.Context) {
	//panic("DO NOT use the function! Use ListOrdersService directly")
	var err error
	var req order_page.ListOrdersReq

	userID, exists := c.Get("user_id")
	if !exists {
		c.String(http.StatusBadRequest, "User ID not found in context")
		return
	}

	req.UserId = userID.(uint32)

	resp, err := o.NewListOrdersService(c).Run(&req)
	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func PlaceOrder(c *gin.Context) {
	var err error
	var req order_page.PlaceOrderReq

	err = c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	resp, err := o.NewPlaceOrderService(c).Run(&req)

	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	c.String(http.StatusOK, "%v", resp)
}

func MarkOrderPaid(c *gin.Context) {
	//panic("DO NOT use the function! Use MarkOrderPaidService directly")
	var err error
	var req order_page.MarkOrderPaidReq

	orderID := c.Query("order_id")
	req.OrderId = orderID

	resp, err := o.NewMarkOrderPaidService(c).Run(&req)
	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
