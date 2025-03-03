package payment

import (
	"github.com/asmile1559/dyshop/app/frontend/biz/service"
	"github.com/asmile1559/dyshop/pb/frontend/payment_page"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Charge(c *gin.Context) {

	var err error
	var req payment_page.ChargeReq

	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "StatusBadRequest",
			"err":     err.Error(),
		})
		return
	}

	resp, err := service.NewChargeService(c).Run(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "StatusBadRequest",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "charge ok!",
		"resp": resp,
	})
}
