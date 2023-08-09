package resp

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

const (
	CodeNoSuchRoute = 1000 // 系统相关错误码
	CodeRequestPara = 1001
	CodeForbidden   = 1002

	CodeNoSuchHost   = 2000 // 业务相关错误码
	CodeNotLogin     = 2001
	CodeInvalidToken = 2002
)

func errorResponse(ctx context.Context, code int, errCode int, err error, msg string, isArray []bool) {
	c, ok := ctx.(*gin.Context)
	if !ok {
		return
	}

	c.Header("X-Request-Id", c.GetString("requestID"))
	if len(isArray) == 0 {
		c.JSON(code, &Response{Code: errCode, Msg: msg, Data: struct{}{}})
	} else {
		c.JSON(code, &ArrayResponse{Code: errCode, Msg: msg, Total: 0, Data: []struct{}{}})
	}

	if err != nil {
		_ = c.Error(errors.WithMessage(err, msg))
	} else {
		_ = c.Error(errors.New(msg))
	}

	c.Set("code", errCode)
	c.Abort()
}

// ErrorRequest 请求参数错误
func ErrorRequest(ctx context.Context, err error, msg string, isArray ...bool) {
	errorResponse(ctx, http.StatusBadRequest, CodeRequestPara, err, msg, isArray)
}

// ErrorRequestWithMsg 请求参数错误
func ErrorRequestWithMsg(ctx context.Context, err error, msg string, isArray ...bool) {
	errorResponse(ctx, http.StatusBadRequest, CodeRequestPara, err, msg, isArray)
}

// ErrorForbidden 无权访问
func ErrorForbidden(ctx context.Context, msg string, isArray ...bool) {
	errorResponse(ctx, http.StatusForbidden, CodeForbidden, nil, msg, isArray)
}

// ErrorInvalidateToken token 无效
func ErrorInvalidateToken(ctx context.Context, msg string, isArray ...bool) {
	errorResponse(ctx, http.StatusForbidden, CodeInvalidToken, nil, msg, isArray)
}

// ErrorNoLogin 用户未登录
func ErrorNoLogin(ctx context.Context, err error, isArray ...bool) {
	errorResponse(ctx, http.StatusUnauthorized, CodeNotLogin, err, "user not login", isArray)
}
