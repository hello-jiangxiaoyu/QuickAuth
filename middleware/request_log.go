package middleware

import (
	"QuickAuth/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type LogFormat struct {
	Ts            string        `json:"ts"`
	Status        int           `json:"status"`
	Method        string        `json:"method"`
	Path          string        `json:"path"`
	Query         string        `json:"query"`
	ClientIp      string        `json:"client_ip"`
	Host          string        `json:"host"`
	Cost          time.Duration `json:"cost"`
	ReqLength     int64         `json:"request_length"`
	BodyBytesSent int           `json:"body_bytes_sent"`
	HttpReferer   string        `json:"http_referer"`
	UserAgent     string        `json:"user_agent"`
	Errors        string        `json:"errors"`
}

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		global.AccessLog.Info("",
			zap.String("ts", time.Now().Format(time.RFC3339)),
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("client_ip", c.ClientIP()),
			zap.String("host", c.Request.Host),
			zap.Duration("cost", time.Since(start)),
			zap.Int64("request_length", c.Request.ContentLength),
			zap.Int("body_bytes_sent", c.Writer.Size()),
			zap.String("http_referer", c.Request.Referer()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		)

		// todo: upload to clickhouse
	}
}
