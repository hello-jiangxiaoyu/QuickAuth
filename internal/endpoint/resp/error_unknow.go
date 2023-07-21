package resp

import (
	"context"
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
	errorResponse(ctx, http.StatusInternalServerError, CodeUnknown, err, respMsg, isArray)
}

// ErrorSqlModify SQL修改失败
func ErrorSqlModify(ctx context.Context, err error, respMsg string, isArray ...bool) {
	errorResponse(ctx, http.StatusInternalServerError, CodeSqlModify, err, respMsg, isArray)
}

// ErrorSelect 数据库查询错误
func ErrorSelect(ctx context.Context, err error, respMsg string, isArray ...bool) {
	errorResponse(ctx, http.StatusInternalServerError, CodeSqlSelect, err, respMsg, isArray)
}

// ErrorNotFound 资源未找到
func ErrorNotFound(ctx context.Context, err error, respMsg string, isArray ...bool) {
	errorResponse(ctx, http.StatusInternalServerError, CodeNotFound, err, respMsg, isArray)
}

func ErrorSaveSession(ctx context.Context, err error, isArray ...bool) {
	errorResponse(ctx, http.StatusInternalServerError, CodeSaveSession, err, "failed to save session", isArray)
}
