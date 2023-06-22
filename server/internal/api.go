package internal

import (
	"QuickAuth/server/model"
	"github.com/gin-gonic/gin"
)

type Api struct {
	c      *gin.Context
	Err    error
	Tenant model.Tenant
}

func NewApi(c *gin.Context) *Api {
	a := &Api{c: c}
	return a
}

func (a *Api) setError(err error) {
	if a.Err == nil {
		a.Err = err
	}
}
func (a *Api) Error() error {
	return a.Err
}
