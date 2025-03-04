package checkout

import (
	"github.com/asmile1559/dyshop/app/frontend/biz/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func Checkout(c *gin.Context) {
	transactionId := c.Query("transaction_id")
	orderId := c.Query("order_id")
	fmt.Println("Transaction ID:", transactionId, ", Order ID:", orderId)

	// 调用 service 获取订单详情
	orderResp, err := getOrderFromBackend(c, orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("获取订单信息失败: %v", err)})
		return
	}

	// 渲染 HTML 页面
	c.HTML(http.StatusOK, "checkout.html", orderResp)
}

// getOrderFromBackend 直接返回 gin.H
func getOrderFromBackend(c *gin.Context, orderId string) (gin.H, error) {
	return service.NewGetOrderWithItemsService(c).Run(orderId)
}
