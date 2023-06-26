package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ServerPanic     = 2000
	CodeUnknown     = 2001
	CodeSqlSelect   = 2002
	CodeSqlModify   = 2002
	CodeNotFound    = 2003
	CodeSaveSession = 2004
)

// ErrorUnknown 未知错误
func ErrorUnknown(c *gin.Context, msg string, isArray ...bool) {
	errorResponse(c, http.StatusInternalServerError, CodeUnknown, msg, isArray)
}

// ErrorSqlModify SQL修改失败
func ErrorSqlModify(c *gin.Context, msg string, isArray ...bool) {
	errorResponse(c, http.StatusInternalServerError, CodeSqlModify, msg, isArray)
}

// ErrorSelect 数据库查询错误
func ErrorSelect(c *gin.Context, msg string, isArray ...bool) {
	errorResponse(c, http.StatusInternalServerError, CodeSqlSelect, msg, isArray)
}

// ErrorNotFound 资源未找到
func ErrorNotFound(c *gin.Context, msg string, isArray ...bool) {
	errorResponse(c, http.StatusInternalServerError, CodeNotFound, msg, isArray)
}

func ErrorSaveSession(c *gin.Context, isArray ...bool) {
	errorResponse(c, http.StatusInternalServerError, CodeSaveSession, "failed to save session", isArray)
}
