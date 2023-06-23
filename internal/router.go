package internal

import (
	"QuickAuth/internal/middleware"
	"QuickAuth/internal/server/controller/oauth"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
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

	oauth.NewOauth2Router(r)
	return r
}
