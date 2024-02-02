package resp

import (
	"context"
	"net/http"
)

// ErrorUnknown 未知错误
func ErrorUnknown(ctx context.Context, err error, respMsg string) {
	errorResponse(ctx, http.StatusInternalServerError, CodeUnknown, err, respMsg)
}

// ErrorNotFound 资源未找到
func ErrorNotFound(ctx context.Context, err error, respMsg string) {
	errorResponse(ctx, http.StatusInternalServerError, CodeNotFound, err, respMsg)
}

func ErrorSaveSession(ctx context.Context, err error) {
	errorResponse(ctx, http.StatusInternalServerError, CodeSaveSession, err, "failed to save session")
}
