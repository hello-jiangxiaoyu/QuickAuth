package utils

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// GetGoID get goroutine id
func GetGoID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

// DeferErr handle defer function err
func DeferErr(errFunc func() error) {
	if err := errFunc(); err != nil {
		fmt.Println("### Defer err: ", err)
	}
}

func DtoFilter[S any, T any](s []S, f func(S) T) []T {
	var l []T
	for _, i := range s {
		l = append(l, f(i))
	}
	return l
}

func GetPanicStackInfo(req string, err any, skip int, fullStack bool) string {
	pwd, _ := os.Getwd()
	pwd = strings.ReplaceAll(pwd, `\`, "/") // handle windows path
	res := fmt.Sprintf("[Recovery] %s panic recovered: %s\n%v", time.Now().Format("2006-01-02 15:04:05"), req, err)
	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		if fullStack || pwd == "" || strings.Contains(file, pwd) { // only about quick auth source files
			res += fmt.Sprintf("\n\t%s:%d %s", file, line, runtime.FuncForPC(pc).Name())
		}
	}
	return res + "\n"
}
