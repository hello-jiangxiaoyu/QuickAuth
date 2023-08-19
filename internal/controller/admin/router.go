package admin

import (
	"QuickAuth/internal/service"
	"QuickAuth/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func AddAdminRoute(svc *service.Service, e *gin.Engine) {
	admin := NewAdminRoute(svc)
	app := e.Group("/api/quick")
	{
		app.GET("/apps", admin.ListApp)
		app.GET("/apps/:appId", admin.GetApp)
		app.POST("/apps", admin.CreateApp)
		app.PUT("/apps/:appId", admin.ModifyApp)
		app.DELETE("/apps/:appId", admin.DeleteApp)

		app.GET("/apps/:appId/secrets", admin.ListAppSecret)
		app.POST("/apps/:appId/secrets", admin.CreateAppSecret)
		app.PUT("/apps/:appId/secrets/:secretId", admin.ModifyAppSecret)
		app.DELETE("/apps/:appId/secrets/:secretId", admin.DeleteAppSecret)

		app.GET("/apps/:appId/tenants", admin.ListTenant)
		app.GET("/apps/:appId/tenants/:tenantId", admin.GetTenant)
		app.POST("/apps/:appId/tenants", admin.CreateTenant)
		app.PUT("/apps/:appId/tenants/:tenantId", admin.ModifyTenant)
		app.DELETE("/apps/:appId/tenants/:tenantId", admin.DeleteTenant)
	}

	tenant := e.Group("/api/quick")
	{
		tenant.GET("/redirect-uri", admin.ListRedirectUri)
		tenant.POST("/redirect-uri", admin.CreateRedirectUri)
		tenant.PUT("/redirect-uri/:uriId", admin.ModifyRedirectUri)
		tenant.DELETE("/redirect-uri/:uri", admin.DeleteRedirectUri)

		tenant.GET("/providers/:providerId", admin.GetProvider)
		tenant.POST("/providers", admin.CreateProvider)
		tenant.PUT("/providers/:providerId", admin.ModifyProvider)
		tenant.DELETE("/providers/:providerId", admin.DeleteProvider)
	}
	e.GET("/api/quick/providers", admin.ListProvider) // 获取当前租户所有第三方登录所需信息

	user := e.Group("/api/quick", middleware.LoginAuth())
	{
		user.GET("/user-pools", admin.ListUserPool)
		user.GET("/user-pools/:poolId", admin.GetUserPool)
		user.POST("/user-pools", admin.CreateUserPool)
		user.PUT("/user-pools/:poolId", admin.ModifyUserPool)
		user.DELETE("/user-pools/:poolId", admin.DeleteUserPool)

		user.GET("/user-pools/:poolId/users", admin.ListUser)
		user.GET("/user-pools/:poolId/users/:userId", admin.GetUser)
		user.POST("/user-pools/:poolId/users", admin.CreateUser)
		user.PUT("/user-pools/:poolId/users/:userId", admin.ModifyUser)
		user.DELETE("/user-pools/:poolId/users/:userId", admin.DeleteUser)
	}
}
