package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StringErrorResponse(c *gin.Context, err error) {
	if err != nil {
		c.String(http.StatusOK, "An error occurred: %v", err)
	}
	return
}
