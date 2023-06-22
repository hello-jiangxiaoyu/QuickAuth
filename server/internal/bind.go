package internal

import "github.com/gin-gonic/gin/binding"

func (a *Api) BindJson(obj any) *Api {
	err := a.c.ShouldBindJSON(obj)
	if err != nil {
		a.setError(err)
	}
	return a
}

func (a *Api) BindUri(obj any) *Api {
	err := a.c.ShouldBindUri(obj)
	if err != nil {
		a.setError(err)
	}
	return a
}

func (a *Api) BindQuery(obj any) *Api {
	err := a.c.BindQuery(obj)
	if err != nil {
		a.setError(err)
	}
	return a
}

func (a *Api) BindForm(obj any) *Api {
	err := a.c.ShouldBindWith(obj, binding.Form)
	if err != nil {
		a.setError(err)
	}
	return a
}
