package resp

import (
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
	Data  []any  `json:"data"`
}

func response(c *gin.Context, code int, errCode uint, msg string, data any) {
	c.JSON(code, &Response{Code: errCode, Msg: msg, Data: data})
	c.Abort()
}
func arrayResponse(c *gin.Context, code int, errCode uint, msg string, total uint, data []any) {
	c.JSON(code, &ArrayResponse{Code: errCode, Msg: msg, Total: total, Data: data})
	c.Abort()
}

func Success(c *gin.Context) {
	response(c, http.StatusOK, CodeSuccess, MsgSuccess, struct{}{})
}
func SuccessWithData(c *gin.Context, data any) {
	response(c, http.StatusOK, CodeSuccess, MsgSuccess, data)
}
func SuccessArray(c *gin.Context, total uint, data []any) {
	arrayResponse(c, http.StatusOK, CodeSuccess, MsgSuccess, total, data)
}

func DoNothing(c *gin.Context, msg string, isArray ...bool) {
	if len(isArray) == 0 {
		response(c, http.StatusAccepted, CodeAccept, msg, struct{}{})
	} else {
		arrayResponse(c, http.StatusAccepted, CodeAccept, msg, 0, []any{})
	}
}
