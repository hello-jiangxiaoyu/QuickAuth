package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
}
func arrayResponse(c *gin.Context, code int, errCode uint, msg string, total uint, data []any) {
	c.JSON(code, &ArrayResponse{Code: errCode, Msg: msg, Total: total, Data: data})
}

func Success(c *gin.Context) {
	response(c, http.StatusOK, 0, "", struct{}{})
}
func SuccessWithData(c *gin.Context, data any) {
	response(c, http.StatusOK, 0, "", data)
}
func SuccessArray(c *gin.Context, total uint, data []any) {
	arrayResponse(c, http.StatusOK, 0, "", total, data)
}
