package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ErrorCodeRequestPara = 1000
	ErrorCodeForbidden   = 1001
	ErrorCodeNoLogin     = 1002
	ErrorNoSuchHost      = 1003
)

// ErrorRequest 请求参数错误
func ErrorRequest(c *gin.Context) {
	response(c, http.StatusBadRequest, ErrorCodeRequestPara, "invalidate request parameters", struct{}{})
	c.Abort()
}
func ErrorRequestArray(c *gin.Context) {
	arrayResponse(c, http.StatusBadRequest, ErrorCodeRequestPara, "invalidate request parameters", 0, []any{})
	c.Abort()
}

// ErrorForbidden 无权访问
func ErrorForbidden(c *gin.Context, msg string) {
	response(c, http.StatusForbidden, ErrorCodeForbidden, msg, struct{}{})
	c.Abort()
}
func ErrorForbiddenArray(c *gin.Context, msg string) {
	arrayResponse(c, http.StatusForbidden, ErrorCodeForbidden, msg, 0, []any{})
	c.Abort()
}

// ErrorNoLogin 用户未登录
func ErrorNoLogin(c *gin.Context) {
	response(c, http.StatusUnauthorized, ErrorCodeNoLogin, "user not login", struct{}{})
	c.Abort()
}
func ErrorNoLoginArray(c *gin.Context) {
	arrayResponse(c, http.StatusUnauthorized, ErrorCodeNoLogin, "user not login", 0, []any{})
	c.Abort()
}
