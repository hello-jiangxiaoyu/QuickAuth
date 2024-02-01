package middleware

import (
	"QuickAuth/pkg/global"
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

func zapLogRequest(c *gin.Context, start time.Time) {
	tenant, _ := getTenant(c)
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
		zap.String("app_id", tenant.AppID),
		zap.Int64("tenant_id", tenant.ID),
		zap.String("user_id", c.GetString("user_id")),
		zap.Int64("pool_id", tenant.UserPool.ID),
		zap.String("tag", tenant.App.Tag),
		zap.String("query", c.Request.URL.RawQuery),
		zap.String("client_ip", c.ClientIP()),
		zap.String("server_ip", getServerIp(c.Request.RemoteAddr)),
		zap.String("host", c.Request.Host),
		zap.Duration("cost", time.Since(start)),
		zap.Int64("request_length", c.Request.ContentLength),
		zap.Int("body_bytes_sent", c.Writer.Size()),
		zap.String("referer", c.Request.Referer()),
		zap.String("proto", c.Request.Proto),
		zap.String("ua", c.Request.UserAgent()),
		zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
	)
}

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		zapLogRequest(c, start)
	}
}

func getServerIp(remoteIp string) string {
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
