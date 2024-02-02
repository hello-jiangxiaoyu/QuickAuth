package resp

import (
	"QuickAuth/pkg/utils"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArrayResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg,omitempty"`
	Total int    `json:"total,omitempty"`
	Data  any    `json:"data"`
}

func response(ctx context.Context, code int, errCode int, err error, msg string, data any, total int) {
	c, ok := ctx.(*gin.Context)
	if !ok {
		return
	}

	c.Header("X-Request-Id", c.GetString("requestID"))

	if data == nil {
		data = []struct{}{}
	}
	c.JSON(code, &ArrayResponse{Code: errCode, Msg: msg, Total: total, Data: data})

	if err != nil {
		_ = c.Error(utils.WithMessage(err, msg))
	} else {
		_ = c.Error(errors.New(msg))
	}
	c.Set("code", errCode)
	c.Abort()
}

func success(ctx context.Context, data any, total int) {
	response(ctx, http.StatusOK, CodeSuccess, nil, MsgSuccess, data, total)
}

func Success(ctx context.Context) {
	success(ctx, struct{}{}, 0)
}
func SuccessWithData(ctx context.Context, data any) {
	success(ctx, data, 0)
}
func SuccessArrayData(ctx context.Context, total int, data any) {
	if total == 0 {
		data = []struct{}{}
	}
	success(ctx, data, total)
}

func DoNothing(ctx context.Context, msg string) {
	response(ctx, http.StatusAccepted, CodeAccept, nil, msg, nil, 0)
}
