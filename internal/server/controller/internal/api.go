package internal

import (
	"github.com/gin-gonic/gin"
)

type Api struct {
	c     *gin.Context
	Error error
}

func NewApi(c *gin.Context) *Api {
	a := &Api{c: c}
	return a
}

func (a *Api) setError(err error) {
	if a.Error == nil {
		a.Error = err
	}
}
