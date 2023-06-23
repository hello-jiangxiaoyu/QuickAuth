package utils

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
)

func GetGoID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

// DeferErr 处理defer返回的错误警告
func DeferErr(errFunc func() error) {
	if err := errFunc(); err != nil {
		fmt.Println("### Defer err: ", err)
	}
}
