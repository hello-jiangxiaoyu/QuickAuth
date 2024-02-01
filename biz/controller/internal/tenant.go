package internal

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/biz/endpoint/resp"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

func (a *Api) SetTenant(t *model.Tenant) *Api {
	if a.c == nil {
		return a.setError(errors.New("gin context should not be nil"))
	}
	value, ok := a.c.Get(resp.Tenant)
	if !ok {
		return a.setError(errors.New("failed to get gin tenant"))
	}

	tenant, ok := value.(model.Tenant)
	if !ok {
		return a.setError(errors.New("failed to convert gin tenant"))
	}

	*t = tenant
	return a
}

func (a *Api) SetUserInfo() *Api {
	if a.c == nil {
		return a.setError(errors.New("gin context should not be nil"))
	}
	value, ok := a.c.Get(resp.UserInfo)
	if !ok {
		return a.setError(errors.New("failed to get gin IDClaims"))
	}
	claim, ok := value.(jwt.MapClaims)
	if !ok {
		return a.setError(errors.New("failed to convert to MapClaims"))
	}
	sub, ok := claim["sub"]
	if !ok {
		return a.setError(errors.New("failed to get sub from MapClaims"))
	}
	a.Sub, ok = sub.(string)
	if !ok {
		return a.setError(errors.New("failed to convert sub to string"))
	}
	a.UserInfo = claim
	return a
}
