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

	c.String(http.StatusOK, "%v", resp)
}
