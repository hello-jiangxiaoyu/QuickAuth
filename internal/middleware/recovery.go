package middleware

import (
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/internal/global"
	"QuickAuth/pkg/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"os"
	"strings"
	"time"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") || strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				req := fmt.Sprintf("requestID:%v method:%s path:%s", c.GetString("requestID"), c.Request.Method, c.Request.URL.Path)
				if global.Config != nil {
					global.Log.Error(utils.GetPanicStackInfo(req, err, 3, global.Config.Log.IsFullStack))
				}
				if brokenPipe {
					_ = c.Error(err.(error))
					c.Abort() // If the connection is dead, we can't write a status to it.
				} else {
					resp.ErrorPanic(c)
				}
				zapLogRequest(c, time.Now())
			}
		}()
		c.Next() // 不能省略
	}
}
