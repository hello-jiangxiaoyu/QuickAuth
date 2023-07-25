package middleware

import (
	"QuickAuth/pkg/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const loggerKey = "log-key"

var Logger *zap.Logger

func NewContext(ctx *gin.Context, fields ...zapcore.Field) {
	ctx.Set(loggerKey, WithContext(ctx).With(fields...))
}

func WithContext(ctx *gin.Context) *zap.Logger {
	if ctx == nil {
		return Logger
	}
	l, ok := ctx.Get(loggerKey)
	if !ok {
		return Logger
	}
	ctxLogger, ok := l.(*zap.Logger)
	if !ok {
		return Logger
	}
	return ctxLogger
}

func traceLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := utils.GetNoLineUUID()
		NewContext(c, zap.String("traceId", traceId))
		NewContext(c, zap.String("request.method", c.Request.Method))
		NewContext(c, zap.String("request.url", c.Request.URL.String()))
		if c.Request.Form == nil {
			if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
				return
			}
		}

		form, err := json.Marshal(c.Request.Form)
		if err != nil {
			return
		}
		NewContext(c, zap.String("request.params", string(form)))
		c.Next()
	}
}
