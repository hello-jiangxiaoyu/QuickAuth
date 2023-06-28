package internal

import (
	"QuickAuth/pkg/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (a *Api) SetTenant(t *model.Tenant) *Api {
	if a.c == nil {
		a.setError(errors.New("gin context should not be nil"))
		return a
	}
	value, ok := a.c.Get("tenant")
	if !ok {
		a.setError(errors.New("failed to get gin tenant"))
		return a
	}

	tenant, ok := value.(model.Tenant)
	if !ok {
		a.setError(errors.New("failed to convert gin tenant"))
		return a
	}

	*t = tenant
	return a
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
