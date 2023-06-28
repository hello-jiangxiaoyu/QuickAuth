package resp

import (
	"QuickAuth/internal/global"
	"context"
	"go.uber.org/zap"
	"net/http"
)

const (
	ServerPanic     = 2000
	CodeUnknown     = 2001
	CodeSqlSelect   = 2002
	CodeSqlModify   = 2002
	CodeNotFound    = 2003
	CodeSaveSession = 2004
)

// ErrorUnknown 未知错误
func ErrorUnknown(ctx context.Context, err error, respMsg string, isArray ...bool) {
	errorResponse(ctx, http.StatusInternalServerError, CodeUnknown, respMsg, isArray)
	global.Log.Error(respMsg, zap.Error(err))
}

// ErrorSqlModify SQL修改失败
func ErrorSqlModify(ctx context.Context, err error, respMsg string, isArray ...bool) {
	errorResponse(ctx, http.StatusInternalServerError, CodeSqlModify, respMsg, isArray)
	global.Log.Error(respMsg, zap.Error(err))
}

// ErrorSelect 数据库查询错误
func ErrorSelect(ctx context.Context, err error, respMsg string, isArray ...bool) {
	errorResponse(ctx, http.StatusInternalServerError, CodeSqlSelect, respMsg, isArray)
	global.Log.Error(respMsg, zap.Error(err))
}

// ErrorNotFound 资源未找到
func ErrorNotFound(ctx context.Context, err error, respMsg string, isArray ...bool) {
	errorResponse(ctx, http.StatusInternalServerError, CodeNotFound, respMsg, isArray)
	global.Log.Error(respMsg, zap.Error(err))
}

func ErrorSaveSession(ctx context.Context, err error, isArray ...bool) {
	errorResponse(ctx, http.StatusInternalServerError, CodeSaveSession, "failed to save session", isArray)
	global.Log.Error("failed to save session: ", zap.Error(err))
}
