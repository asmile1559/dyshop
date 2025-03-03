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
		_checkout.POST("/checkout", append(_checkout1Mw(), ch.Checkout)...)

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
