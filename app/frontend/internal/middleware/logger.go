package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		// 记录请求日志
		logrus.WithFields(logrus.Fields{
			"path":    path,
			"query":   raw,
			"ip":      c.ClientIP(),
			"method":  c.Request.Method,
			"status":  c.Writer.Status(),
			"latency": time.Since(start),
		}).Info("access log")
	}
}
