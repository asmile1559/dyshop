package order

import (
	o "github.com/asmile1559/dyshop/app/frontend/biz/service"
	"github.com/asmile1559/dyshop/pb/frontend/order_page"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOrders(c *gin.Context) {
	panic("DO NOT use the function! Use ListOrdersService directly")
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
	panic("DO NOT use the function! Use MarkOrderPaidService directly")
}
