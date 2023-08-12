package resp

import (
	"context"
	"net/http"
	"strings"
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
	if err != nil && strings.HasPrefix(err.Error(), "ERROR: duplicate key value violates unique constraint") {
		errorResponse(ctx, http.StatusInternalServerError, CodeSqlCreateDuplicate, err, "Duplicate field name", isArray)
	} else {
		errorResponse(ctx, http.StatusInternalServerError, CodeSqlCreate, err, respMsg, isArray)
	}
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
