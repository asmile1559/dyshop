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
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	resp, err := service.NewChargeService(c).Run(&req)
	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
		return
	}

	c.String(http.StatusOK, "%v", resp)
}
