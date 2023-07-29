package resp

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

const (
	CodeRequestPara = 1000
	CodeForbidden   = 1001
	CodeNoLogin     = 1002
	CodeNoSuchHost  = 1003
	CodeNoSuchRoute = 1003
)

func errorResponse(ctx context.Context, code int, errCode int, err error, msg string, isArray []bool) {
	c, ok := ctx.(*gin.Context)
	if !ok {
		return
	}

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

// ErrorNoLogin 用户未登录
func ErrorNoLogin(ctx context.Context, isArray ...bool) {
	errorResponse(ctx, http.StatusUnauthorized, CodeNoLogin, nil, "user not login", isArray)
}
