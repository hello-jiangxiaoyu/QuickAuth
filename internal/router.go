package internal

import (
	"QuickAuth/internal/controller"
	"QuickAuth/internal/global"
	"QuickAuth/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
	r := gin.Default()
	// r := gin.New()
	// gin.SetMode(gin.ReleaseMode)
	r.Use(middleware.Recovery(), middleware.RequestLog(), cors.Default())
	cookieSecret := []byte("QuickAuth")
	store := cookie.NewStore(cookieSecret)
	store.Options(sessions.Options{
		MaxAge: 60 * 60 * 24 * 3,
		Path:   "/",
	})
	r.Use(sessions.Sessions("QuickAuth", store))

	controller.NewOauth2Router(&global.Repository{Logger: global.Log, DB: global.DB, Config: global.Config}, r)
	return r
}
