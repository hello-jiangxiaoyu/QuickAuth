package internal

import (
	"QuickAuth/internal/controller"
	"QuickAuth/internal/global"
	"QuickAuth/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
	r := gin.Default()
	// r := gin.New()
	// gin.SetMode(gin.ReleaseMode)
	r.Use(middleware.Recovery(), cors.Default())
	r.Use(middleware.RequestLog(), middleware.GenerateRequestID())
	r.Use(middleware.TenantHost())
	controller.NewOauth2Router(&global.Repository{Logger: global.Log, DB: global.DB, Config: global.Config}, r)
	return r
}
