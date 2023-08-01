package middleware

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
)

func GenerateRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = strconv.FormatInt(rand.Int63(), 10)
		}
		c.Set("requestID", requestID)
	}
}
