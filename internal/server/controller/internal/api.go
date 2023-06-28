package internal

import (
	"github.com/gin-gonic/gin"
)

type Api struct {
	c     *gin.Context
	Error error
}

func (a *Api) SetCtx(c *gin.Context) *Api {
	a.c = c
	return a
}

func (a *Api) setError(err error) {
	if a.Error == nil {
		a.Error = err
	}
}
