package admin

import (
	"QuickAuth/internal/service"
	"QuickAuth/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// AddAdminRoute admin接口
func AddAdminRoute(svc *service.Service, e *gin.Engine) {
	admin := NewAdminRoute(svc)
	app := e.Group("/api/quick", middleware.LoginAuth())
	{
		// app应用管理
		app.GET("/apps", admin.ListApp)
		app.GET("/apps/:appId", admin.GetApp)
		app.POST("/apps", admin.CreateApp)
		app.PUT("/apps/:appId", admin.ModifyApp)
		app.DELETE("/apps/:appId", admin.DeleteApp)

		// m2m 密钥和权限
		app.GET("/apps/:appId/secrets", admin.ListAppSecret)
		app.POST("/apps/:appId/secrets", admin.CreateAppSecret)
		app.PUT("/apps/:appId/secrets/:secretId", admin.ModifyAppSecret)
		app.DELETE("/apps/:appId/secrets/:secretId", admin.DeleteAppSecret)

		// 租户管理
		app.GET("/apps/:appId/tenants", admin.ListTenant)
		app.GET("/apps/:appId/tenants/:tenantId", admin.GetTenant)
		app.POST("/apps/:appId/tenants", admin.CreateTenant)
		app.PUT("/apps/:appId/tenants/:tenantId", admin.ModifyTenant)
		app.DELETE("/apps/:appId/tenants/:tenantId", admin.DeleteTenant)

		// 重定向uri管理
		app.GET("/redirect-uri", admin.ListRedirectUri)
		app.POST("/redirect-uri", admin.CreateRedirectUri)
		app.PUT("/redirect-uri/:uriId", admin.ModifyRedirectUri)
		app.DELETE("/redirect-uri/:uri", admin.DeleteRedirectUri)

		// provider第三方id系统
		app.GET("/providers/:providerId", admin.GetProvider)
		app.POST("/providers", admin.CreateProvider)
		app.PUT("/providers/:providerId", admin.ModifyProvider)
		app.DELETE("/providers/:providerId", admin.DeleteProvider)
	}
	e.GET("/api/quick/providers", admin.ListProvider) // 获取当前租户所有第三方登录所需信息

	// 用户管理
	user := e.Group("/api/quick", middleware.LoginAuth())
	{
		// 用户池
		user.GET("/user-pools", admin.ListUserPool)
		user.GET("/user-pools/:poolId", admin.GetUserPool)
		user.POST("/user-pools", admin.CreateUserPool)
		user.PUT("/user-pools/:poolId", admin.ModifyUserPool)
		user.DELETE("/user-pools/:poolId", admin.DeleteUserPool)

		// 实体用户
		user.GET("/user-pools/:poolId/users", admin.ListUser)
		user.GET("/user-pools/:poolId/users/:userId", admin.GetUser)
		user.POST("/user-pools/:poolId/users", admin.CreateUser)
		user.PUT("/user-pools/:poolId/users/:userId", admin.ModifyUser)
		user.DELETE("/user-pools/:poolId/users/:userId", admin.DeleteUser)
	}
}
