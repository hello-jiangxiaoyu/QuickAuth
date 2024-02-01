package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Api struct {
	c        *gin.Context
	UserInfo jwt.MapClaims
	Sub      string
	Error    error
}

func New(c *gin.Context) *Api {
	return &Api{
		c: c,
	}
}

func (a *Api) SetCtx(c *gin.Context) *Api {
	a.c = c
	return a
}

func (a *Api) setError(err error) *Api {
	if a.Error == nil {
		a.Error = err
	}
	return a
}
