package oauth

import (
	"QuickAuth/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func AddOauth2Route(e *gin.Engine) {
	o := NewOAuth2Route()
	r := e.Group("/api/quick")
	{
		r.GET("/.well-known/openid-configuration", o.GetOIDC) // OIDC信息
		r.GET("/.well-known/jwks.json", o.ListJwks)           // jwk签名公钥

		r.POST("/login", o.Login)                                          // 账号密码登录到当前平台
		r.GET("/login/providers/:providerId", o.ProviderLogin)             // 跳转到第三方登录
		r.GET("/login/providers/:providerId/callback", o.ProviderCallback) // 第三方登录回调，登录到当前平台（QuickAuth）
		r.GET("/oauth2/auth", middleware.LoginAuth(), o.GetAuthCode)       // 授权第三方
		r.POST("/oauth2/token", middleware.M2mAuth(), o.GetToken)          // 服务端根据code获取token

		r.POST("/register", o.Register)                            // 注册
		r.GET("/logout", o.Logout)                                 // 登出
		r.GET("/me/profile", middleware.LoginAuth(), o.GetProfile) // 客户端获取当前用户信息
	}
}
