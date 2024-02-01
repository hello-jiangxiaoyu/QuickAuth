package resp

import (
	"context"
	"net/http"
)

// ErrorUnknown 未知错误
func ErrorUnknown(ctx context.Context, err error, respMsg string, isArray ...bool) {
	errorResponse(ctx, http.StatusInternalServerError, CodeUnknown, err, respMsg, isArray)
}

// ErrorNotFound 资源未找到
func ErrorNotFound(ctx context.Context, err error, respMsg string, isArray ...bool) {
	errorResponse(ctx, http.StatusInternalServerError, CodeNotFound, err, respMsg, isArray)
}

func ErrorSaveSession(ctx context.Context, err error, isArray ...bool) {
	errorResponse(ctx, http.StatusInternalServerError, CodeSaveSession, err, "failed to save session", isArray)
}
