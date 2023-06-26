package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodeRequestPara = 1000
	CodeForbidden   = 1001
	CodeNoLogin     = 1002
	CodeNoSuchHost  = 1003
)

func errorResponse(c *gin.Context, code int, errCode uint, msg string, isArray []bool) {
	if len(isArray) == 0 {
		response(c, code, errCode, msg, struct{}{})
	} else {
		arrayResponse(c, code, errCode, msg, 0, []any{})
	}
}

// ErrorRequest 请求参数错误
func ErrorRequest(c *gin.Context, isArray ...bool) {
	errorResponse(c, http.StatusBadRequest, CodeRequestPara, "invalidate request parameters", isArray)
}

// ErrorRequestWithMsg 请求参数错误
func ErrorRequestWithMsg(c *gin.Context, msg string, isArray ...bool) {
	errorResponse(c, http.StatusBadRequest, CodeRequestPara, msg, isArray)
}

// ErrorForbidden 无权访问
func ErrorForbidden(c *gin.Context, msg string, isArray ...bool) {
	errorResponse(c, http.StatusForbidden, CodeForbidden, msg, isArray)
}

// ErrorNoLogin 用户未登录
func ErrorNoLogin(c *gin.Context, isArray ...bool) {
	errorResponse(c, http.StatusUnauthorized, CodeNoLogin, "user not login", isArray)
}
