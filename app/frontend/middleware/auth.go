package middleware

import (
	"errors"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
	"github.com/sirupsen/logrus"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken, err := c.Cookie("token")
		logrus.Debug(authToken)
		if errors.Is(err, http.ErrNoCookie) {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "No token provided",
			})
			c.Abort()
			return
		}

		token := strings.Split(authToken, " ")
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
		logrus.WithFields(logrus.Fields{
			"token":  token[1],
			"method": method,
			"uri":    uri,
		}).Debug("auth middleware token")
		resp, err := rpcclient.AuthClient.VerifyTokenByRPC(c, &pbauth.VerifyTokenReq{
			Token:  token[1],
			Method: method,
			Uri:    uri,
		})
		logrus.WithFields(logrus.Fields{
			"success": resp.Res,
			"userid": resp.UserId,
		}).Debug("verify token resp")
		if err != nil {
			logrus.WithError(err).Debug("AuthClient.VerifyTokenByRPC err")
		}

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
