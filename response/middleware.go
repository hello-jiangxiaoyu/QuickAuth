package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorPanic(c *gin.Context) {
	response(c, http.StatusInternalServerError, ErrorServerPanic, "server panic", struct{}{})
	c.Abort()
}

func ErrorHost(c *gin.Context) {
	response(c, http.StatusForbidden, ErrorNoSuchHost, "no such host", struct{}{})
	c.Abort()
}
