package order

import (
	"fmt"
	o "github.com/asmile1559/dyshop/app/frontend/biz/handler/order"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(e *gin.Engine) {
	root := e.Group("/", _rootMw()...)
	{
		_order := root.Group("/order", _orderMw()...)
		// TODO: order front page
		// GET /example/order
		// 获得订单
		_order.GET("/", func(c *gin.Context) {
			orderId := c.Query("order_id")
			resp := gin.H{
				"PageRouter": PageRouter,
				"UserInfo": gin.H{
					"Name": "lixiaoming",
				},
				"AddressInfo": gin.H{
					"Default": "1",
					"Addresses": []gin.H{
						{
							"AddressId":   "1",
							"Recipient":   "张三李四",
							"Phone":       "12345678901",
							"Province":    "中国",
							"City":        "北京市",
							"District":    "海淀区",
							"Street":      "知春路",
							"FullAddress": "北京北京市海淀区知春路甲48号抖音视界",
						},
						{
							"AddressId":   "2",
							"Recipient":   "张三",
							"Phone":       "12345678901",
							"Province":    "广东省",
							"City":        "深圳市",
							"District":    "南山区",
							"Street":      "海天二路",
							"FullAddress": "广东省深圳市南山区海天二路33号腾讯滨海大厦"},
					},
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
				"OrderPrice":      "58.59",
				"OrderPostage":    "10.00",
				"OrderDiscount":   "0",
				"OrderFinalPrice": "68.59",
			}

			fmt.Println(orderId)
			fmt.Println(resp)
			c.HTML(http.StatusOK, "order.html", resp)
		})
		// POST /example/order/submit
		// 用户提交订单
		_order.POST("/place", append(_placeOrderMw(), o.PlaceOrder)...)
		_order.POST("/submit", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "no user id error",
				})
				return
			}
			req := struct {
				OrderId         string  `json:"order_id"`
				Address         gin.H   `json:"address"`
				Products        []gin.H `json:"products"`
				Discount        string  `json:"discount"`
				OrderPrice      string  `json:"order_price"`
				OrderPostage    string  `json:"order_postage"`
				OrderFinalPrice string  `json:"order_final_price"`
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
				"message": "submit order ok",
				"resp": gin.H{
					"user_id":        userId,
					"order_id":       req.OrderId,
					"transaction_id": 123,
				}, //TODO transaction_id改为rpc调用后的数据
			})
		}) /**/

		// POST /example/order/cancel
		// 用户取消订单
		_order.POST("/cancel", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "no user id error",
				})
				return
			}
			req := struct {
				OrderId string `json:"order_id"`
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
					"user_id":  userId,
					"order_id": req.OrderId,
				},
			})
		})

	}
}
