package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorPanic(c *gin.Context) {
	response(c, http.StatusInternalServerError, ServerPanic, "server panic", struct{}{})
	c.Abort()
}

func ErrorHost(c *gin.Context) {
	response(c, http.StatusForbidden, CodeNoSuchHost, "no such host", struct{}{})
	c.Abort()
}
