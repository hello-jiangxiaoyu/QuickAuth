package oauth

import (
	"QuickAuth/internal/service"
	"QuickAuth/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func AddOauth2Route(svc *service.Service, e *gin.Engine) {
	o := NewOAuth2Route(svc)
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
	}
}
