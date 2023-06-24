package utils

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"runtime"
	"strconv"
	"strings"
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

func GetHostWithScheme(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	if s := c.Request.Header.Get("X-Forwarded-Proto"); s != "" {
		scheme = s
	}

	return fmt.Sprintf("%s://%s", scheme, c.Request.Host)
}
func isIpAddress(host string) bool {
	hostWithoutPort, _, err := net.SplitHostPort(host)
	if err != nil {
		hostWithoutPort = host
	}

	ip := net.ParseIP(hostWithoutPort)
	return ip != nil
}

func GetUrlByHost(host string) string {
	protocol := "https"
	if strings.HasPrefix(host, "localhost") || isIpAddress(host) {
		protocol = "http"
	}
	return fmt.Sprintf("%s://%s", protocol, host)
}
