package internal

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
)

func (a *Api) BindJson(obj any) *Api {
	if a.c == nil {
		a.setError(errors.New("gin context should not be nil"))
		return a
	}
	if err := a.c.ShouldBindJSON(obj); err != nil {
		a.setError(err)
	}
	return a
}

func (a *Api) BindUri(obj any) *Api {
	if a.c == nil {
		a.setError(errors.New("gin context should not be nil"))
		return a
	}
	if err := a.c.ShouldBindUri(obj); err != nil { // bind path param
		a.setError(err)
	}
	return a
}

func (a *Api) BindUriAndJson(obj any) *Api {
	if a.c == nil {
		a.setError(errors.New("gin context should not be nil"))
		return a
	}
	if err := a.c.ShouldBindUri(obj); err != nil { // bind path param
		a.setError(err)
		return a
	}
	if err := a.c.ShouldBindJSON(obj); err != nil {
		a.setError(err)
	}
	return a
}

func (a *Api) BindQuery(obj any) *Api {
	if a.c == nil {
		a.setError(errors.New("gin context should not be nil"))
		return a
	}
	if err := a.c.BindQuery(obj); err != nil {
		a.setError(err)
	}
	return a
}

func (a *Api) BindForm(obj any) *Api {
	if a.c == nil {
		a.setError(errors.New("gin context should not be nil"))
		return a
	}
	if err := a.c.ShouldBindWith(obj, binding.Form); err != nil {
		a.setError(err)
	}
	return a
}
