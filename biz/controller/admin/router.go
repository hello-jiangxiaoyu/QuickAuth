package admin

import (
	"QuickAuth/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// AddAdminRoute admin接口
func AddAdminRoute(e *gin.RouterGroup) {
	app := e.Group("", middleware.LoginAuth())
	{
		// app应用管理
		app.GET("/apps", ListApp)
		app.GET("/apps/:appId", GetApp)
		app.POST("/apps", CreateApp)
		app.PUT("/apps/:appId", ModifyApp)
		app.DELETE("/apps/:appId", DeleteApp)

		// m2m 密钥和权限
		app.GET("/apps/:appId/secrets", ListAppSecret)
		app.POST("/apps/:appId/secrets", CreateAppSecret)
		app.PUT("/apps/:appId/secrets/:secretId", ModifyAppSecret)
		app.DELETE("/apps/:appId/secrets/:secretId", DeleteAppSecret)

		// 租户管理
		app.GET("/apps/:appId/tenants", ListTenant)
		app.GET("/apps/:appId/tenants/:tenantId", GetTenant)
		app.POST("/apps/:appId/tenants", CreateTenant)
		app.PUT("/apps/:appId/tenants/:tenantId", ModifyTenant)
		app.DELETE("/apps/:appId/tenants/:tenantId", DeleteTenant)

		// 重定向uri管理
		app.GET("/redirect-uri", ListRedirectUri)
		app.POST("/redirect-uri", CreateRedirectUri)
		app.PUT("/redirect-uri/:uriId", ModifyRedirectUri)
		app.DELETE("/redirect-uri/:uri", DeleteRedirectUri)

		// provider第三方id系统
		app.GET("/providers/:providerId", GetProvider)
		app.POST("/providers", CreateProvider)
		app.PUT("/providers/:providerId", ModifyProvider)
		app.DELETE("/providers/:providerId", DeleteProvider)
	}
	e.GET("/providers", ListLoginProviderInfo) // 获取当前租户所有第三方登录所需信息

	// 用户管理
	user := e.Group("", middleware.LoginAuth())
	{
		// 用户池
		user.GET("/user-pools", ListUserPool)
		user.GET("/user-pools/:poolId", GetUserPool)
		user.POST("/user-pools", CreateUserPool)
		user.PUT("/user-pools/:poolId", ModifyUserPool)
		user.DELETE("/user-pools/:poolId", DeleteUserPool)

		// 实体用户
		user.GET("/user-pools/:poolId/users", ListUser)
		user.GET("/user-pools/:poolId/users/:userId", GetUser)
		user.POST("/user-pools/:poolId/users", CreateUser)
		user.PUT("/user-pools/:poolId/users/:userId", ModifyUser)
		user.DELETE("/user-pools/:poolId/users/:userId", DeleteUser)
	}
}
