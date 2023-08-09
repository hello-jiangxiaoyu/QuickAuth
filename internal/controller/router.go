package controller

import (
	_ "QuickAuth/docs"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/internal/global"
	"QuickAuth/internal/middleware"
	"QuickAuth/internal/service"
	"strings"

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
		r.GET("/.well-known/openid-configuration", o.getOIDC)      // OIDC信息
		r.GET("/.well-known/jwks.json", o.getJwks)                 // jwk签名公钥
		r.GET("/oauth2/auth", o.getAuthCode)                       // 登录授权
		r.POST("/oauth2/token", o.getToken)                        // token获取
		r.GET("/me/profile", middleware.LoginAuth(), o.getProfile) // 获取当前用户信息

		r.POST("/login", o.login)                               // 账号密码登录
		r.POST("/register", o.register)                         // 注册
		r.GET("/logout", o.logout)                              // 登出
		r.GET("/login/providers/:provider", o.providerCallback) // 第三方登录回调
		r.GET("/providers", o.listProvider)                     // 获取当前租户所有第三方登录所需信息
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

		app.GET("/apps/:appId/tenants", o.listTenant)
		app.GET("/apps/:appId/tenants/:tenantId", o.getTenant)
		app.POST("/apps/:appId/tenants", o.createTenant)
		app.PUT("/apps/:appId/tenants/:tenantId", o.modifyTenant)
		app.DELETE("/apps/:appId/tenants/:tenantId", o.deleteTenant)
	}

	tenant := e.Group("/api/quick")
	{
		tenant.GET("/redirect-uri", o.listRedirectUri)
		tenant.POST("/redirect-uri", o.createRedirectUri)
		tenant.PUT("/redirect-uri/:uriId", o.modifyRedirectUri)
		tenant.DELETE("/redirect-uri/:uri", o.deleteRedirectUri)

		tenant.GET("/providers/:providerId", o.getProvider)
		tenant.POST("/providers", o.createProvider)
		tenant.PUT("/providers/:providerId", o.modifyProvider)
		tenant.DELETE("/providers/:providerId", o.deleteProvider)
	}

	user := e.Group("/api/quick", middleware.LoginAuth())
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
	AddWebRoutes(e)
}

func AddWebRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) { // 首页
		c.Redirect(http.StatusMovedPermanently, "/applications/")
	})

	r.Use(middleware.StaticWebFile()) // 其他静态资源
	r.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api/quick") {
			resp.ErrorNoRoute(c)
			return
		}
		c.File("web/out/404/index.html")
	})
}
