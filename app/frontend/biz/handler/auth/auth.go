package auth

import (
	"github.com/gin-gonic/gin"
)

func DeliverToken(c *gin.Context) {

	panic("DO NOT use the function! Use DeliverTokenService directly")
	
	//var err error
	//var req auth_page.DeliverTokenReq
	//
	//err = c.Bind(&req)
	//if err != nil {
	//	c.String(http.StatusOK, "An error occurred: %v", err)
	//	return
	//}
	//
	//resp, err := service.NewDeliverTokenService(c).Run(&req)
	//
	//if err != nil {
	//	c.String(http.StatusOK, "An error occurred: %v", err)
	//	return
	//}
	//
	//c.String(http.StatusOK, "DeliverToken ok! your token is: %v", resp.Token)
}

func VerifyToken(c *gin.Context) {

	panic("DO NOT use the function! Use VerifyTokenService directly")

	//var err error
	//var req auth_page.VerifyTokenReq
	//
	//err = c.Bind(&req)
	//if err != nil {
	//	c.String(http.StatusOK, "An error occurred: %v", err)
	//	return
	//}
	//
	//resp, err := service.NewVerifyService(c).Run(&req)
	//
	//if err != nil {
	//	c.String(http.StatusOK, "An error occurred: %v", err)
	//	return
	//}
	//
	//c.String(http.StatusOK, "VerifyToken ok! your token verify resulT is: %v", resp.Res)
}
