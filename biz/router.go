package internal

import (
	"QuickAuth/biz/controller"
	"QuickAuth/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.Use(middleware.Recovery(), cors.Default())
	r.Use(middleware.RequestLog(), middleware.GenerateRequestID())
	r.Use(middleware.TenantHost())
	controller.NewRouter(r)
	return r
}
