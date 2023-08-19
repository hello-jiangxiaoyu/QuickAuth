package controller

import (
	_ "QuickAuth/docs"
	"QuickAuth/internal/controller/admin"
	"QuickAuth/internal/controller/oauth"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/internal/service"
	"QuickAuth/pkg/global"
	"QuickAuth/pkg/middleware"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewOauth2Router(repo *global.Repository, e *gin.Engine) {
	svc := service.NewService(repo)
	a := admin.AddAdminRoute(svc, e)
	o := oauth.NewOAuth2Route(svc)

	r := e.Group("/api/quick")
	{
		r.GET("/.well-known/openid-configuration", o.GetOIDC)      // OIDC信息
		r.GET("/.well-known/jwks.json", o.ListJwks)                // jwk签名公钥
		r.GET("/oauth2/auth", o.GetAuthCode)                       // 登录授权
		r.POST("/oauth2/token", o.GetToken)                        // token获取
		r.GET("/me/profile", middleware.LoginAuth(), o.GetProfile) // 获取当前用户信息

		r.POST("/login", o.Login)                               // 账号密码登录
		r.POST("/register", o.Register)                         // 注册
		r.GET("/logout", o.Logout)                              // 登出
		r.GET("/login/providers/:provider", o.ProviderCallback) // 第三方登录回调
		r.GET("/providers", a.ListProvider)                     // 获取当前租户所有第三方登录所需信息
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
