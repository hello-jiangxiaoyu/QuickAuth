package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ErrorCodeUnknown = 2000
	ErrorCodeSql     = 2001
	ErrorNotFound    = 2002
)

func ErrorUnknown(c *gin.Context, msg string) {
	response(c, http.StatusInternalServerError, ErrorCodeUnknown, msg, struct{}{})
}
func ErrorUnknownArray(c *gin.Context, msg string) {
	arrayResponse(c, http.StatusInternalServerError, ErrorCodeUnknown, msg, 0, []any{})
}
