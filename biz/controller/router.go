package controller

import (
	"QuickAuth/biz/controller/abac"
	"QuickAuth/biz/controller/admin"
	_ "QuickAuth/biz/controller/internal/docs"
	"QuickAuth/biz/controller/oauth"
	"QuickAuth/biz/controller/rg"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/pkg/middleware"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewRouter(e *gin.Engine) {
	api := e.Group("/api/quick", middleware.TenantHost())
	{
		admin.AddAdminRoute(api)
		oauth.AddOauth2Route(api)
		abac.AddIamRouter(api)
		rg.AddResourceGroupRoutes(api)
	}

	AddWebRoutes(e)

	e.GET("/api/quick/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	e.GET("/api/quick/health", func(c *gin.Context) { c.String(http.StatusOK, "ok") }) // 健康探测
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
