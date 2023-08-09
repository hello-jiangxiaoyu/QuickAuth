package internal

import (
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/pkg/model"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

func (a *Api) SetTenant(t *model.Tenant) *Api {
	if a.c == nil {
		a.setError(errors.New("gin context should not be nil"))
		return a
	}
	value, ok := a.c.Get(resp.Tenant)
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

func (a *Api) SetUserInfo() *Api {
	if a.c == nil {
		a.setError(errors.New("gin context should not be nil"))
		return a
	}
	value, ok := a.c.Get(resp.UserInfo)
	if !ok {
		a.setError(errors.New("failed to get gin IDClaims"))
		return a
	}
	claim, ok := value.(jwt.MapClaims)
	if !ok {
		a.setError(errors.New("failed to convert to MapClaims"))
		return a
	}
	sub, ok := claim["sub"]
	if !ok {
		a.setError(errors.New("failed to get sub from MapClaims"))
	}
	a.Sub, ok = sub.(string)
	if !ok {
		a.setError(errors.New("failed to convert sub to string"))
		return a
	}
	a.UserInfo = claim
	return a
}
