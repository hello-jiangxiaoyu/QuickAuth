package controller

import (
	"QuickAuth/internal/global"
	"QuickAuth/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewOauth2Router(repo *global.Repository, e *gin.Engine) {
	svc := service.NewService(repo)
	o := NewOAuth2Api(svc)
	r := e.Group("/api/quick")
	{
		r.GET("/.well-known/openid-configuration", o.getOIDC)   // OIDC信息
		r.GET("/.well-known/jwks.json", o.getJwks)              // jwk签名公钥
		r.POST("/login", o.login)                               // 账号密码登录
		r.GET("/login/providers/:provider", o.providerCallback) // 第三方登录回调
		r.GET("/providers", o.listProvider)                     // 获取当前租户所有第三方登录所需信息
		r.GET("/logout", o.logout)                              // 登出
		r.GET("/oauth2/auth", o.getAuthCode)                    // 登录授权
		r.POST("/oauth2/token", o.getToken)                     // token获取
	}

	client := e.Group("/api/quick/clients")
	{
		client.GET("", o.listClient)
		client.GET("/:clientId", o.getClient)
		client.POST("", o.createClient)
		client.PUT("/:clientId", o.modifyClient)
		client.DELETE("/:clientId", o.deleteClient)

		client.GET("/:clientId/secrets", o.listClientSecret)
		client.POST("/:clientId/secrets", o.createClientSecret)
		client.DELETE("/:clientId/secrets/:secretId", o.deleteClientSecret)

		client.GET("/:clientId/redirect-uri", o.listRedirectUri)
		client.POST("/:clientId/redirect-uri", o.createRedirectUri)
		client.PUT("/:clientId/redirect-uri/:uriId", o.modifyRedirectUri)
		client.DELETE("/:clientId/redirect-uri/:uriId", o.deleteClientSecret)
	}

	tenant := e.Group("/api/quick")
	{
		tenant.GET("/:clientId/tenants", o.listTenant)
		tenant.GET("/:clientId/tenants/:tenantId", o.getTenant)
		tenant.POST("/:clientId/tenants", o.createTenant)
		tenant.PUT("/:clientId/tenants/:tenantId", o.modifyTenant)
		tenant.DELETE("/:clientId/tenants/:tenantId", o.deleteTenant)
	}
	provider := e.Group("/api/quick/providers") // 通过host区分租户
	{
		provider.GET("/:providerId", o.getProvider)
		provider.POST("", o.createProvider)
		provider.PUT("/:providerId", o.modifyProvider)
		provider.DELETE("/:providerId", o.deleteProvider)
	}

	user := e.Group("/api/quick/user-pools")
	{
		user.GET("/:poolId")
	}

	e.GET("/api/quick/health", func(c *gin.Context) { c.String(http.StatusOK, "ok") }) // 健康探测
}