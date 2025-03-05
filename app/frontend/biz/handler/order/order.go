package order

import (
	"net/http"
	"strconv"

	o "github.com/asmile1559/dyshop/app/frontend/biz/service"
	"github.com/asmile1559/dyshop/pb/frontend/order_page"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ListOrders(c *gin.Context) {
	var req order_page.ListOrderReq

	orderId := c.Query("order_id")
	oid, _ := strconv.Atoi(orderId)
	if oid == 0 {
		resp, err := o.NewListOrdersService(c).Run(&req)
		if err != nil {
			c.String(http.StatusOK, "An error occurred: %v", err)
			return
		}
		c.HTML(http.StatusOK, "order.html", resp)
		return
	}
	// get by order id
	resp, err := o.NewGetOrderService(c).Run(&order_page.GetOrderReq{OrderId: uint32(oid)})
	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}
	logrus.Debug("=>", resp)
	c.HTML(http.StatusOK, "order.html", resp)
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
