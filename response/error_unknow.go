package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ErrorServerPanic = 2000
	ErrorCodeUnknown = 2001
	ErrorCodeSql     = 2002
	ErrorNotFound    = 2003
)

func ErrorUnknown(c *gin.Context, msg string) {
	response(c, http.StatusInternalServerError, ErrorCodeUnknown, msg, struct{}{})
	c.Abort()
}

func ErrorUnknownArray(c *gin.Context, msg string) {
	arrayResponse(c, http.StatusInternalServerError, ErrorCodeUnknown, msg, 0, []any{})
	c.Abort()
}
