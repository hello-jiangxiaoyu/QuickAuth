package internal

import (
	_ "QuickAuth/docs"
	"QuickAuth/internal/global"
	"QuickAuth/internal/middleware"
	"QuickAuth/internal/server/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GetServer() *gin.Engine {
	r := gin.New()
	gin.Recovery()
	r.Use(middleware.Recovery(), middleware.RequestLog(), cors.Default())
	cookieSecret := []byte("QuickAuth")
	store := cookie.NewStore(cookieSecret)
	store.Options(sessions.Options{
		MaxAge: 60 * 60 * 24 * 3,
		Path:   "/",
	})
	r.Use(sessions.Sessions("QixinAuth", store))

	r.GET("/quick/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	controller.NewOauth2Router(&global.Repository{Logger: global.Log, DB: global.DB, Config: global.Config}, r)
	return r
}
