package initial

import (
	"QuickAuth/middleware"
	"QuickAuth/server/controller"
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

	controller.NewOauth2Router(r)
	return r
}
