package middleware

import (
	"time"

	"github.com/asmile1559/dyshop/utils/logx"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.StandardLogger()
	defaultFormatter := logx.DefaultFormatter
	defaultFormatter.Role = "USER"
	log.SetFormatter(defaultFormatter)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		// 记录请求日志
		log.WithFields(logrus.Fields{
			"path":    path,
			"query":   raw,
			"ip":      c.ClientIP(),
			"method":  c.Request.Method,
			"status":  c.Writer.Status(),
			"latency": time.Since(start),
		}).Info("access log")
	}
}
