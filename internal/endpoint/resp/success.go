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

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, &Response{Code: CodeSuccess, Msg: MsgSuccess, Data: struct{}{}})
	c.Abort()
}
func SuccessWithData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, &Response{Code: CodeSuccess, Msg: MsgSuccess, Data: data})
	c.Abort()
}
func SuccessArray(c *gin.Context, total uint, data []any) {
	c.JSON(http.StatusOK, &ArrayResponse{Code: CodeSuccess, Msg: MsgSuccess, Total: total, Data: data})
	c.Abort()
}

func DoNothing(c *gin.Context, msg string, isArray ...bool) {
	if len(isArray) == 0 {
		c.JSON(http.StatusAccepted, &Response{Code: CodeAccept, Msg: msg, Data: struct{}{}})
	} else {
		c.JSON(http.StatusAccepted, &ArrayResponse{Code: CodeAccept, Msg: msg, Total: 0, Data: []any{}})
	}
	c.Abort()
}
