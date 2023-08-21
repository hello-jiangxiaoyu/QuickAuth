package resp

import (
	"context"
	"net/http"
)

func errorResponse(ctx context.Context, code int, errCode int, err error, msg string, isArray []bool) {
	if err != nil { // todo: 临时开发使用，将错误信息直接返回给前端，项目稳定后需要删除
		msg += ": " + err.Error()
	}
	response(ctx, code, errCode, err, msg, nil, 0, isArray)
}

// ErrorRequest 请求参数错误
func ErrorRequest(ctx context.Context, err error, isArray ...bool) {
	errorResponse(ctx, http.StatusBadRequest, CodeRequestPara, err, "invalid request param", isArray)
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
func ErrorInvalidateToken(ctx context.Context, isArray ...bool) {
	errorResponse(ctx, http.StatusForbidden, CodeInvalidToken, nil, "invalidated token", isArray)
}

// ErrorNoLogin 用户未登录
func ErrorNoLogin(ctx context.Context, err error, isArray ...bool) {
	errorResponse(ctx, http.StatusUnauthorized, CodeNotLogin, err, "user not login", isArray)
}
