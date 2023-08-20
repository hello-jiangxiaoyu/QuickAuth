package controller

import (
	_ "QuickAuth/docs"
	"QuickAuth/internal/controller/admin"
	"QuickAuth/internal/controller/iam"
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

func NewRouter(repo *global.Repository, e *gin.Engine) {
	svc := service.NewService(repo)

	admin.AddAdminRoute(svc, e)
	oauth.AddOauth2Route(svc, e)
	iam.AddIamRouter(svc, e)
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
