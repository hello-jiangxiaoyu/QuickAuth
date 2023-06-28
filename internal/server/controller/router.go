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
		r.GET("/.well-known/openid-configuration", o.getOIDC) // OIDC信息
		r.GET("/.well-known/jwks.json", o.getJwks)            // jwk签名公钥
		r.POST("/login", o.login)                             // 账号密码登录
		r.GET("/logout", o.logout)                            // 登出
		r.GET("/oauth2/auth", o.getAuthCode)                  // 登录授权
		r.POST("/oauth2/token", o.getToken)                   // token获取

		r.GET("/login/provider-info", o.listProviders)         // 获取当前租户所有第三方登录所需信息
		r.GET("/login/provider/:provider", o.providerCallback) // 第三方登录回调
	}
	e.GET("/api/quick/health", func(c *gin.Context) { c.String(http.StatusOK, "ok") }) // 健康探测
}
