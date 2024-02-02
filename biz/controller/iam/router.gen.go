package iam

import (
	"github.com/gin-gonic/gin"
)

func AddIamRouter(e *gin.RouterGroup) {
	// 资源管理
	resource := e.Group("")
	{
		resource.GET("/resources", ListResources)
		resource.GET("/resources/:resourceId", GetResource)
		resource.POST("/resources", CreateResource)
		resource.PUT("/resources/:resourceId", UpdateResource)
		resource.DELETE("/resources/:resourceId", DeleteResource)
	}

	// 资源属性管理
	resourceAttribute := e.Group("/resources/:resourceId")
	{
		// 资源下的节点
		resourceAttribute.GET("/nodes", ListResourceNodes)
		resourceAttribute.GET("/nodes/:nodeId", GetResourceNode)
		resourceAttribute.POST("/nodes", CreateResourceNode)
		resourceAttribute.PUT("/nodes/:nodeId", UpdateResourceNode)
		resourceAttribute.DELETE("/nodes/:nodeId", DeleteResourceNode)

		// 资源下的角色
		resourceAttribute.GET("/roles", ListResourceRoles)
		resourceAttribute.GET("/roles/:roleId", GetResourceRole)
		resourceAttribute.POST("/roles", CreateResourceRole)
		resourceAttribute.PUT("/roles/:roleId", UpdateResourceRole)
		resourceAttribute.DELETE("/roles/:roleId", DeleteResourceRole)

		// 资源下的操作
		resourceAttribute.GET("/operations", ListResourceOperations)
		resourceAttribute.GET("/operations/:operationId", GetResourceOperation)
		resourceAttribute.POST("/operations", CreateResourceOperation)
		resourceAttribute.PUT("/operations/:operationId", UpdateResourceOperation)
		resourceAttribute.DELETE("/operations/:operationId", DeleteResourceOperation)
	}

	// 授权管理
	auth := e.Group("/resources/:resourceId")
	{
		// 角色的权限管理
		auth.GET("/roles/:roleId/operations", ListResourceRoleOperations)
		auth.POST("/roles/:roleId/operations", CreateResourceRoleOperation)
		auth.DELETE("/roles/:roleId/operations/:operationId", DeleteResourceRoleOperation)

		// json资源用户的角色管理
		auth.GET("/users/:userId/roles", ListResourceUserRoles)
		auth.POST("/users/:userId/roles", CreateResourceUserRole)
		auth.PUT("/users/:userId/roles/:roleId", UpdateResourceUserRole)
		auth.DELETE("/users/:userId/roles/:roleId", DeleteResourceUserRole)

		auth.GET("/json/users/:userId/roles", ListResourceJSONUserRoles)
		auth.POST("/json/users/:userId/roles", CreateResourceJSONUserRole)
		auth.PUT("/json/users/:userId/roles/:roleId", UpdateResourceJSONUserRole)
		auth.DELETE("/json/users/:userId/roles/:roleId", DeleteResourceJSONUserRole)
	}

	// 鉴权
	iamAuth := e.Group("/resources/:resourceId")
	{
		iamAuth.GET("/nodes/:nodeId/operations/:operationId", IsOperationAllow) // 针对某个资源的操作，判断是否允许
		iamAuth.GET("/json/operations/:operationId", IsJSONOperationAllow)      // 针对某个json资源的操作，判断是否允许
		iamAuth.GET("/operations/:operationId/parents/:parentId", ListResourceOperationNodes)
		iamAuth.GET("/operations/:operationId/json", ListJSONResourceOperationNodes)
	}
}
