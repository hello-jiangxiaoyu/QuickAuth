package resp

import (
	"QuickAuth/internal/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
		c.JSON(code, &Response{Code: errCode, Msg: msg, Data: struct{}{}})
	} else {
		c.JSON(code, &ArrayResponse{Code: errCode, Msg: msg, Total: 0, Data: []any{}})
	}
	c.Abort()
}

// ErrorRequest 请求参数错误
func ErrorRequest(c *gin.Context, err error, msg string, isArray ...bool) {
	errorResponse(c, http.StatusBadRequest, CodeRequestPara, msg, isArray)
	global.Log.Error(msg, zap.Error(err))
}

// ErrorRequestWithMsg 请求参数错误
func ErrorRequestWithMsg(c *gin.Context, err error, msg string, isArray ...bool) {
	errorResponse(c, http.StatusBadRequest, CodeRequestPara, msg, isArray)
	global.Log.Error(msg, zap.Error(err))
}

// ErrorForbidden 无权访问
func ErrorForbidden(c *gin.Context, msg string, isArray ...bool) {
	errorResponse(c, http.StatusForbidden, CodeForbidden, msg, isArray)
}

// ErrorNoLogin 用户未登录
func ErrorNoLogin(c *gin.Context, isArray ...bool) {
	errorResponse(c, http.StatusUnauthorized, CodeNoLogin, "user not login", isArray)
}
