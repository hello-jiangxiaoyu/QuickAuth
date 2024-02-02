package resp

import (
	"context"
	"net/http"
)

func ErrorPanic(ctx context.Context) {
	errorResponse(ctx, http.StatusInternalServerError, CodeServerPanic, nil, "server panic")
}

func ErrorHost(ctx context.Context) {
	errorResponse(ctx, http.StatusForbidden, CodeNoSuchHost, nil, "no such host")
}

func ErrorNoRoute(ctx context.Context) {
	errorResponse(ctx, http.StatusNotFound, CodeNoSuchRoute, nil, "no such route")
}
