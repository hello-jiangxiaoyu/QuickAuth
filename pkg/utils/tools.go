package utils

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
	"strconv"
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

// WithMessage err和msg有一个不为空则返回error
func WithMessage(err error, msg string) error {
	if err == nil && msg == "" {
		return nil
	} else if err == nil {
		return errors.New(msg)
	} else if msg == "" {
		return err
	}

	return errors.New(msg + ": " + err.Error())
}
