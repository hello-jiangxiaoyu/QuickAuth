package resp

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodeSuccess = 0
	CodeAccept  = 0
	MsgSuccess  = ""
)

type Response struct {
	Code uint   `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type ArrayResponse struct {
	Code  uint   `json:"code"`
	Msg   string `json:"msg"`
	Total uint   `json:"total"`
	Data  any    `json:"data"`
}

func success(ctx context.Context, data any, total uint, isArray bool) {
	c, ok := ctx.(*gin.Context)
	if !ok {
		return
	}

	if !isArray {
		c.JSON(http.StatusOK, &Response{Code: CodeSuccess, Msg: MsgSuccess, Data: data})
	} else {
		c.JSON(http.StatusOK, &ArrayResponse{Code: CodeSuccess, Msg: MsgSuccess, Total: total, Data: data})
	}
	c.Abort()
}

func Success(ctx context.Context) {
	success(ctx, struct{}{}, 0, false)
}
func SuccessWithData(ctx context.Context, data any) {
	success(ctx, data, 0, false)
}
func SuccessArray(ctx context.Context, total uint, data []any) {
	success(ctx, data, total, true)
}

func DoNothing(ctx context.Context, msg string, isArray ...bool) {
	c, ok := ctx.(*gin.Context)
	if !ok {
		return
	}

	if len(isArray) == 0 {
		c.JSON(http.StatusAccepted, &Response{Code: CodeAccept, Msg: msg, Data: struct{}{}})
		c.Abort()
	} else {
		c.JSON(http.StatusAccepted, &ArrayResponse{Code: CodeAccept, Msg: msg, Total: 0, Data: []struct{}{}})
		c.Abort()
	}
}
