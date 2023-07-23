package resp

import (
	"context"
	"net/http"
)

const (
	ServerPanic     = 2000
	CodeUnknown     = 2001
	CodeNotFound    = 2002
	CodeSaveSession = 2003

	CodeSqlSelect = 3000
	CodeSqlModify = 3001
	CodeSqlCreate = 3002
	CodeSqlDelete = 3003
)

// ErrorUnknown 未知错误
func ErrorUnknown(ctx context.Context, err error, respMsg string, isArray ...bool) {
	errorResponse(ctx, http.StatusInternalServerError, CodeUnknown, err, respMsg, isArray)
}

// ErrorSqlModify SQL修改失败
func ErrorSqlModify(ctx context.Context, err error, respMsg string, isArray ...bool) {
	errorResponse(ctx, http.StatusInternalServerError, CodeSqlModify, err, respMsg, isArray)
}

func ErrorSqlCreate(ctx context.Context, err error, respMsg string, isArray ...bool) {
	errorResponse(ctx, http.StatusInternalServerError, CodeSqlCreate, err, respMsg, isArray)
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
