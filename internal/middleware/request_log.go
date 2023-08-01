package middleware

import (
	"QuickAuth/internal/global"
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

		appId := c.GetString("appId")
		tenantId := c.GetInt("tenantId")
		userId := c.GetString("userId")
		poolId := c.GetInt("poolId")
		tag := c.GetString("tag")
		code := c.GetInt("code")
		requestID := c.GetString("requestID")
		global.AccessLog.Info("",
			zap.String("ts", time.Now().Format(time.RFC3339)),
			zap.String("request_id", requestID),
			zap.Int("status", c.Writer.Status()),
			zap.Int("error_code", code),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("full_path", c.FullPath()),
			zap.String("app_id", appId),
			zap.Int("tenant_id", tenantId),
			zap.String("user_id", userId),
			zap.Int("pool_id", poolId),
			zap.String("tag", tag),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("client_ip", c.ClientIP()),
			zap.String("server_ip", GetServerIp(c.Request.RemoteAddr)),
			zap.String("host", c.Request.Host),
			zap.Duration("cost", time.Since(start)),
			zap.Int64("request_length", c.Request.ContentLength),
			zap.Int("body_bytes_sent", c.Writer.Size()),
			zap.String("referer", c.Request.Referer()),
			zap.String("proto", c.Request.Proto),
			zap.String("ua", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		)

		// todo: upload to clickhouse
	}
}

func GetServerIp(remoteIp string) string {
	if len(remoteIp) > 0 && remoteIp[0] == ':' {
		return remoteIp
	}
	for i := len(remoteIp) - 1; i >= 0; i-- {
		if remoteIp[i] == ':' {
			remoteIp = remoteIp[:i]
			break
		}
	}
	return remoteIp
}
