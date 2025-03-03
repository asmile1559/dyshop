package checkout

import (
	"github.com/asmile1559/dyshop/app/frontend/biz/service"
	"github.com/asmile1559/dyshop/pb/frontend/checkout_page"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Checkout(c *gin.Context) {

	var err error
	var req checkout_page.CheckoutReq

	//transactionId := c.Query("transaction_id")
	orderId := c.Query("order_id")


	err = c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	resp, err := service.NewCheckoutService(c).Run(&req)
	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	resp = gin.H{
		"PageRouter": PageRouter,
		"UserInfo": gin.H{
			"Name": "lixiaoming",
		},
		"OrderId": orderId,
		"Address": gin.H{
			"Recipient":   "张三李四",
			"Phone":       "12345678901",
			"Province":    "北京",
			"City":        "北京市",
			"District":    "海淀区",
			"Street":      "知春路",
			"FullAddress": "北京北京市海淀区知春路甲48号抖音视界",
		},
		"Products": []gin.H{
			{
				"ProductId":   "1",
				"ProductImg":  "/static/src/product/bearcookie.webp",
				"ProductName": "超级无敌好吃的小熊饼干",
				"ProductSpec": gin.H{
					"Name":  "500g装",
					"Price": "18.80",
				},
				"Quantity": "2",
				"Currency": "CNY",
				"Postage":  "10.00",
			},
			{
				"ItemId":      "2",
				"productId":   "2",
				"ProductImg":  "/static/src/product/bearsweet.webp",
				"ProductName": "超级无敌好吃的小熊软糖值得品尝大力推荐",
				"ProductSpec": gin.H{
					"Name":  "9分软",
					"Price": "20.99",
				},
				"Quantity": "1",
				"Postage":  "0",
			},
		},
		"OrderQuantity":   "3",
		"OrderPostage":    "10.00",
		"OrderPrice":      "58.59",
		"OrderFinalPrice": "68.59",
	}

	c.HTML(http.StatusOK, "checkout.html", resp)
}
