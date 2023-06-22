package internal

import (
	"QuickAuth/server/model"
	"errors"
)

func (a *Api) SetTenant() *Api {
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

	a.Tenant = tenant
	return a
}
