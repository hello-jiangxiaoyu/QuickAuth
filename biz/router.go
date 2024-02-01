package internal

import (
	"QuickAuth/biz/controller"
	"QuickAuth/pkg/global"
	"QuickAuth/pkg/middleware"
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
	controller.NewRouter(&global.Repository{Logger: global.Log, DB: global.Db(), Config: global.Config}, r)
	return r
}
