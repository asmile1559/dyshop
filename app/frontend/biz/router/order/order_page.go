package order

import (
	"fmt"
	"net/http"

	o "github.com/asmile1559/dyshop/app/frontend/biz/handler/order"
	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine) {
	root := e.Group("/", _rootMw()...)
	{
		_order := root.Group("/order", _orderMw()...)
		// TODO: order front page
		// GET /example/order
		// 获得订单
		_order.GET("/", o.ListOrders)
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
