package controller

import (
	"QuickAuth/internal/global"
	"QuickAuth/internal/server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewOauth2Router(repo *global.Repository, e *gin.Engine) {
	svc := service.NewService(repo)
	o := NewOAuth2Api(svc)
	r := e.Group("/api/quick")
	{
		r.GET("/.well-known/openid-configuration", o.getOIDC)
		r.GET("/.well-known/jwks.json", o.getJwks)
		r.POST("/login", o.login)
		r.GET("/oauth2/auth", o.getAuthCode)
		r.POST("/oauth2/token", o.getToken)
	}
	e.GET("/api/quick/health", func(c *gin.Context) { c.String(http.StatusOK, "ok") })
}
