package resp

import (
	"context"
	"net/http"
)

func errorResponse(ctx context.Context, code int, errCode int, err error, msg string, isArray []bool) {
	response(ctx, code, errCode, err, msg, nil, 0, isArray)
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
