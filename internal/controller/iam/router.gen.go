package iam

import (
	"QuickAuth/internal/service"
	"github.com/gin-gonic/gin"
)

func AddResourceRouter(svc *service.Service, e *gin.Engine) {
	resourceCtl := NewResourceController(svc)

	// 资源管理
	resource := e.Group("/api/quick")
	{
		resource.GET("/resources", resourceCtl.ListResources)
		resource.GET("/resources/:resourceId", resourceCtl.GetResource)
		resource.POST("/resources", resourceCtl.CreateResource)
		resource.PUT("/resources/:resourceId", resourceCtl.UpdateResource)
		resource.DELETE("/resources/:resourceId", resourceCtl.DeleteResource)
	}

	// 资源属性管理
	resourceAttribute := e.Group("/api/quick/resources/:resourceId")
	{
		// 资源下的节点
		resourceAttribute.GET("/nodes", resourceCtl.ListResourceNodes)
		resourceAttribute.GET("/nodes/:nodeId", resourceCtl.GetResourceNode)
		resourceAttribute.POST("/nodes", resourceCtl.CreateResourceNode)
		resourceAttribute.PUT("/nodes/:nodeId", resourceCtl.UpdateResourceNode)
		resourceAttribute.DELETE("/nodes/:nodeId", resourceCtl.DeleteResourceNode)

		// 资源下的角色
		resourceAttribute.GET("/roles", resourceCtl.ListResourceRoles)
		resourceAttribute.GET("/roles/:roleId", resourceCtl.GetResourceRole)
		resourceAttribute.POST("/roles", resourceCtl.CreateResourceRole)
		resourceAttribute.PUT("/roles/:roleId", resourceCtl.UpdateResourceRole)
		resourceAttribute.DELETE("/roles/:roleId", resourceCtl.DeleteResourceRole)

		// 资源下的操作
		resourceAttribute.GET("/operations", resourceCtl.ListResourceOperations)
		resourceAttribute.GET("/operations/:operationId", resourceCtl.GetResourceOperation)
		resourceAttribute.POST("/operations", resourceCtl.CreateResourceOperation)
		resourceAttribute.PUT("/operations/:operationId", resourceCtl.UpdateResourceOperation)
		resourceAttribute.DELETE("/operations/:operationId", resourceCtl.DeleteResourceOperation)
	}

	// 授权管理
	auth := e.Group("/api/quick/resources/:resourceId")
	{
		// 角色的权限管理
		auth.GET("/roles/:roleId/operations", resourceCtl.ListResourceRoleOperations)
		auth.POST("/roles/:roleId/operations", resourceCtl.CreateResourceRoleOperation)
		auth.DELETE("/roles/:roleId/operations/:operationId", resourceCtl.DeleteResourceRoleOperation)

		// json资源用户的角色管理
		auth.GET("/users/:userId/roles", resourceCtl.ListResourceUserRoles)
		auth.POST("/users/:userId/roles", resourceCtl.CreateResourceUserRole)
		auth.PUT("/users/:userId/roles/:roleId", resourceCtl.UpdateResourceUserRole)
		auth.DELETE("/users/:userId/roles/:roleId", resourceCtl.DeleteResourceUserRole)

		auth.GET("/json/users/:userId/roles", resourceCtl.ListResourceJsonUserRoles)
		auth.POST("/json/users/:userId/roles", resourceCtl.CreateResourceJsonUserRole)
		auth.PUT("/json/users/:userId/roles/:roleId", resourceCtl.UpdateResourceJsonUserRole)
		auth.DELETE("/json/users/:userId/roles/:roleId", resourceCtl.DeleteResourceJsonUserRole)

		// 获取拥有某个权限的所有节点
		auth.GET("/operations/:operationId/nodes", resourceCtl.ListResourceOperationNodes)
	}
}
