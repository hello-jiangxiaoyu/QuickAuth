package oauth

import (
	"QuickAuth/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func AddOauth2Route(e *gin.RouterGroup) {
	r := e.Group("")
	{
		r.GET("/.well-known/openid-configuration", GetOIDC) // OIDC信息
		r.GET("/.well-known/jwks.json", ListJwks)           // jwk签名公钥

		r.POST("/login", Login)                                          // 账号密码登录到当前平台
		r.GET("/login/providers/:providerId", ProviderLogin)             // 跳转到第三方登录
		r.GET("/login/providers/:providerId/callback", ProviderCallback) // 第三方登录回调，登录到当前平台（QuickAuth）
		r.GET("/oauth2/auth", middleware.LoginAuth(), GetAuthCode)       // 授权第三方
		r.POST("/oauth2/token", middleware.M2mAuth(), GetToken)          // 服务端根据code获取token

		r.POST("/register", Register)                            // 注册
		r.GET("/logout", Logout)                                 // 登出
		r.GET("/me/profile", middleware.LoginAuth(), GetProfile) // 客户端获取当前用户信息
	}
}
