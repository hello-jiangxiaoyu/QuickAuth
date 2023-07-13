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

	app := e.Group("/api/quick/apps")
	{
		app.GET("", o.listApp)
		app.GET("/:appId", o.getApp)
		app.POST("", o.createApp)
		app.PUT("/:appId", o.modifyApp)
		app.DELETE("/:appId", o.deleteApp)

		app.GET("/:appId/secrets", o.listAppSecret)
		app.POST("/:appId/secrets", o.createAppSecret)
		app.DELETE("/:appId/secrets/:secretId", o.deleteAppSecret)

		app.GET("/:appId/redirect-uri", o.listRedirectUri)
		app.POST("/:appId/redirect-uri", o.createRedirectUri)
		app.PUT("/:appId/redirect-uri/:uriId", o.modifyRedirectUri)
		app.DELETE("/:appId/redirect-uri/:uri", o.deleteRedirectUri)
	}

	tenant := e.Group("/api/quick/apps/:appId/tenants")
	{
		tenant.GET("", o.listTenant)
		tenant.GET("/:tenantId", o.getTenant)
		tenant.POST("", o.createTenant)
		tenant.PUT("/:tenantId", o.modifyTenant)
		tenant.DELETE("/:tenantId", o.deleteTenant)
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
		user.GET("", o.listUserPool)
		user.GET("/:poolId", o.getUserPool)
		user.POST("", o.createUserPool)
		user.PUT("/:poolId", o.modifyUserPool)
		user.DELETE("/:poolId", o.deleteUserPool)

		user.GET("/:poolId/users", o.listUser)
		user.GET("/:poolId/users/:userId", o.getUser)
		user.POST("/:poolId/users", o.createUser)
		user.PUT("/:poolId/users/:userId", o.modifyUser)
		user.DELETE("/:poolId/users/:userId", o.deleteUser)
	}

	e.GET("/api/quick/health", func(c *gin.Context) { c.String(http.StatusOK, "ok") }) // 健康探测
}
