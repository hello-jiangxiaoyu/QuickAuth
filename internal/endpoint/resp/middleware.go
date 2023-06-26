package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorPanic(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, &Response{Code: ServerPanic, Msg: "server panic", Data: struct{}{}})
	c.Abort()
}

func ErrorHost(c *gin.Context) {
	c.JSON(http.StatusForbidden, &Response{Code: CodeNoSuchHost, Msg: "no such host", Data: struct{}{}})
	c.Abort()
}
