package middleware

import (
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/pkg/global"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"os"
	"runtime"
	"strings"
	"time"
)

func GetPanicStackInfo(msg string, err any, skip int, fullStack bool) string {
	pwd, _ := os.Getwd()
	pwd = strings.ReplaceAll(pwd, `\`, "/") // handle windows path
	res := fmt.Sprintf("\n[Recovery] panic recovered: %s\n[Error] %v", msg, err)
	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if !strings.Contains(file, "github.com") && !strings.Contains(file, "gorm.io/") &&
			!strings.Contains(file, "net/http") && !strings.Contains(file, "runtime/") {
			res += fmt.Sprintf("\n\t%s:%d %s", file, line, runtime.FuncForPC(pc).Name())
		}
	}
	return res + "\n"
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
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
					global.Log.Error(GetPanicStackInfo(req, err, 3, global.Config.Log.IsFullStack))
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
