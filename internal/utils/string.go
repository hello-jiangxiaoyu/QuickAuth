package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetHostWithScheme(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	if s := c.Request.Header.Get("X-Forwarded-Proto"); s != "" {
		scheme = s
	}

	return fmt.Sprintf("%s://%s", scheme, c.Request.Host)
}
