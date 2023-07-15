package controller

import (
	_ "QuickAuth/docs"
	"QuickAuth/internal/global"
	"QuickAuth/internal/middleware"
	"QuickAuth/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
		r.GET("/me/profile", o.getProfile)                      // 获取当前用户信息
	}

	app := e.Group("/api/quick")
	{
		app.GET("/apps", o.listApp)
		app.GET("/apps/:appId", o.getApp)
		app.POST("/apps", o.createApp)
		app.PUT("/apps/:appId", o.modifyApp)
		app.DELETE("/apps/:appId", o.deleteApp)

		app.GET("/apps/:appId/secrets", o.listAppSecret)
		app.POST("/apps/:appId/secrets", o.createAppSecret)
		app.PUT("/apps/:appId/secrets/:secretId", o.modifyAppSecret)
		app.DELETE("/apps/:appId/secrets/:secretId", o.deleteAppSecret)

	}

	tenant := e.Group("/api/quick/apps/:appId", middleware.TenantHost())
	{
		tenant.GET("/tenants", o.listTenant)
		tenant.GET("/tenants/current", o.getTenant)
		tenant.POST("/tenants", o.createTenant)
		tenant.PUT("/tenants", o.modifyTenant)
		tenant.DELETE("/tenants", o.deleteTenant)

		app.GET("/redirect-uri", o.listRedirectUri)
		app.POST("/redirect-uri", o.createRedirectUri)
		app.PUT("/redirect-uri/:uriId", o.modifyRedirectUri)
		app.DELETE("/redirect-uri/:uri", o.deleteRedirectUri)
	}

	provider := e.Group("/api/quick", middleware.TenantHost())
	{
		provider.GET("/providers/:providerId", o.getProvider)
		provider.POST("/providers", o.createProvider)
		provider.PUT("/providers/:providerId", o.modifyProvider)
		provider.DELETE("/providers/:providerId", o.deleteProvider)
	}

	user := e.Group("/api/quick")
	{
		user.GET("/user-pools", o.listUserPool)
		user.GET("/user-pools/:poolId", o.getUserPool)
		user.POST("/user-pools", o.createUserPool)
		user.PUT("/user-pools/:poolId", o.modifyUserPool)
		user.DELETE("/user-pools/:poolId", o.deleteUserPool)

		user.GET("/user-pools/:poolId/users", o.listUser)
		user.GET("/user-pools/:poolId/users/:userId", o.getUser)
		user.POST("/user-pools/:poolId/users", o.createUser)
		user.PUT("/user-pools/:poolId/users/:userId", o.modifyUser)
		user.DELETE("/user-pools/:poolId/users/:userId", o.deleteUser)
	}

	e.GET("/api/quick/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	e.GET("/api/quick/health", func(c *gin.Context) { c.String(http.StatusOK, "ok") }) // 健康探测
}
