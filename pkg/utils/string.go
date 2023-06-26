package utils

import (
	"fmt"
	"github.com/google/uuid"
	"net"
	"strings"
)

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

func GetNoLineUUID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}
