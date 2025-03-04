package checkout

import (
	ch "github.com/asmile1559/dyshop/app/frontend/biz/handler/checkout"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func Register(e *gin.Engine) {
	_root := e.Group("/", _rootMw()...)
	{
		_checkout := _root.Group("/checkout", _checkoutMw()...)
		// TODO: checkout front page
		//_checkout.GET("/")
		// GET /example/checkout?transaction_id=&order_id=
		// 进行结算
		_checkout.GET("/", func(c *gin.Context) {
			transactionId := c.Query("transaction_id")
			orderId := c.Query("order_id")

			fmt.Println(transactionId, ",", orderId)

			resp := gin.H{
				"PageRouter": ch.PageRouter,
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
		})

		// GET /example/checkout/cancel
		// 取消结算
		_checkout.POST("/cancel", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "no user id error",
				})
				return
			}
			req := struct {
				OrderId       string `json:"order_id"`
				TransactionId string `json:"transaction_id"`
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

			fmt.Println(req)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "del account ok",
				"resp": gin.H{
					"user_id":        userId,
					"order_id":       req.OrderId,
					"transaction_id": req.TransactionId,
				},
			})
		})
		
	}
}
