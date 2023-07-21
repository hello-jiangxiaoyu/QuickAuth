package resp

import (
	"context"
	"net/http"
)

func ErrorPanic(ctx context.Context) {
	errorResponse(ctx, http.StatusInternalServerError, ServerPanic, nil, "server panic", nil)
}

func ErrorHost(ctx context.Context) {
	errorResponse(ctx, http.StatusForbidden, CodeNoSuchHost, nil, "no such host", nil)
}

func ErrorNoRoute(ctx context.Context) {
	errorResponse(ctx, http.StatusNotFound, CodeNoSuchRoute, nil, "no such route", nil)
}
