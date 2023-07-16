package resp

import (
	"context"
	"net/http"
)

func ErrorPanic(ctx context.Context) {
	errorResponse(ctx, http.StatusInternalServerError, ServerPanic, "server panic", nil)
}

func ErrorHost(ctx context.Context) {
	errorResponse(ctx, http.StatusForbidden, CodeNoSuchHost, "no such host", nil)
}

func ErrorNoRoute(ctx context.Context) {
	errorResponse(ctx, http.StatusNotFound, CodeNoSuchRoute, "no such route", nil)
}
