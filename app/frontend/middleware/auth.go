package middleware

import (
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"

	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Authorization header is required",
			})
			c.Abort()
			return
		}

		token := strings.Split(authHeader, " ")
		if token[0] != "Bearer" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Invalid Authorization header format. Expected 'Bearer <token>'",
			})
			c.Abort()
			return
		}

		method := c.Request.Method
		uri := c.Request.RequestURI
		resp, _ := rpcclient.AuthClient.VerifyTokenByRPC(c, &pbauth.VerifyTokenReq{
			Token:  token[1],
			Method: method,
			Uri:    uri,
		})

		if !resp.GetRes() {
			var message string
			code := int(resp.GetCode())
			switch code {
			case http.StatusUnauthorized:
				message = "Invalid or expired token"
			case http.StatusInternalServerError:
				message = "Something went wrong. Please try again later"
			case http.StatusForbidden:
				message = "Access denied"
			}
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
			c.Abort()
			return
		}
		
		c.Set("user_id", resp.GetUserId())
		c.Next()
	}
}
