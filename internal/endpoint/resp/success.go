package resp

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type ArrayResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Total int    `json:"total"`
	Data  any    `json:"data"`
}

func response(ctx context.Context, code int, errCode int, err error, msg string, data any, total int, isArray []bool) {
	c, ok := ctx.(*gin.Context)
	if !ok {
		return
	}

	c.Header("X-Request-Id", c.GetString("requestID"))
	if len(isArray) == 0 || !isArray[0] {
		if data == nil {
			data = struct{}{}
		}
		c.JSON(code, &Response{Code: errCode, Msg: msg, Data: data})
	} else {
		if data == nil {
			data = []struct{}{}
		}
		c.JSON(code, &ArrayResponse{Code: errCode, Msg: msg, Total: total, Data: data})
	}

	if err != nil {
		_ = c.Error(errors.WithMessage(err, msg))
	} else {
		_ = c.Error(errors.New(msg))
	}
	c.Set("code", errCode)
	c.Abort()
}

func success(ctx context.Context, data any, total int, isArray ...bool) {
	response(ctx, http.StatusOK, CodeSuccess, nil, MsgSuccess, data, total, isArray)
}

func Success(ctx context.Context) {
	success(ctx, struct{}{}, 0, false)
}
func SuccessWithData(ctx context.Context, data any) {
	success(ctx, data, 0, false)
}
func SuccessArrayData(ctx context.Context, total int, data any) {
	if total == 0 {
		data = []struct{}{}
	}
	success(ctx, data, total, true)
}

func DoNothing(ctx context.Context, msg string, isArray ...bool) {
	response(ctx, http.StatusAccepted, CodeAccept, nil, msg, nil, 0, isArray)
}
