package controller

import (
	"QuickAuth/internal/controller/iam"
	"github.com/gin-gonic/gin"
)

func AddResourceRoleRouter(e *gin.Engine) {
	roleCtl := iam.NewResourceRoleController()
	role := e.Group("")
	{
		role.GET("/api/quick/resources/:resourceId/roles", roleCtl.ListResourceRoles)
		role.GET("/api/quick/resources/:resourceId/roles/:roleId", roleCtl.GetResourceRole)
		role.POST("/api/quick/resources/:resourceId/roles", roleCtl.CreateResourceRole)
		role.PUT("/api/quick/resources/:resourceId/roles/:roleId", roleCtl.UpdateResourceRole)
		role.DELETE("/api/quick/resources/:resourceId/roles/:roleId", roleCtl.DeleteResourceRole)
	}
}

func AddResourceRoleOperationRouter(e *gin.Engine) {
	operationCtl := iam.NewResourceRoleOperationController()
	operation := e.Group("")
	{
		operation.GET("/api/quick/resources/:resourceId/roles/:roleId/operations", operationCtl.ListResourceRoleOperations)
		operation.GET("/api/quick/resources/:resourceId/roles/:roleId/operations/:operationId", operationCtl.GetResourceRoleOperation)
		operation.POST("/api/quick/resources/:resourceId/roles/:roleId/operations", operationCtl.CreateResourceRoleOperation)
		operation.PUT("/api/quick/resources/:resourceId/roles/:roleId/operations/:operationId", operationCtl.UpdateResourceRoleOperation)
		operation.DELETE("/api/quick/resources/:resourceId/roles/:roleId/operations/:operationId", operationCtl.DeleteResourceRoleOperation)
	}
}

func AddResourceUserRoleRouter(e *gin.Engine) {
	roleCtl := iam.NewResourceUserRoleController()
	role := e.Group("")
	{
		role.GET("/api/quick/resources/:resourceId/users/:userId/roles", roleCtl.ListResourceUserRoles)
		role.GET("/api/quick/resources/:resourceId/users/:userId/roles/:roleId", roleCtl.GetResourceUserRole)
		role.POST("/api/quick/resources/:resourceId/users/:userId/roles", roleCtl.CreateResourceUserRole)
		role.PUT("/api/quick/resources/:resourceId/users/:userId/roles/:roleId", roleCtl.UpdateResourceUserRole)
		role.DELETE("/api/quick/resources/:resourceId/users/:userId/roles/:roleId", roleCtl.DeleteResourceUserRole)
	}
}

func AddResourceJsonUserRoleRouter(e *gin.Engine) {
	roleCtl := iam.NewResourceJsonUserRoleController()
	role := e.Group("")
	{
		role.GET("/api/quick/resources/:resourceId/json/users/:userId/roles", roleCtl.ListResourceJsonUserRoles)
		role.GET("/api/quick/resources/:resourceId/json/users/:userId/roles/:roleId", roleCtl.GetResourceJsonUserRole)
		role.POST("/api/quick/resources/:resourceId/json/users/:userId/roles", roleCtl.CreateResourceJsonUserRole)
		role.PUT("/api/quick/resources/:resourceId/json/users/:userId/roles/:roleId", roleCtl.UpdateResourceJsonUserRole)
		role.DELETE("/api/quick/resources/:resourceId/json/users/:userId/roles/:roleId", roleCtl.DeleteResourceJsonUserRole)
	}
}

func AddResourceOperationNodeRouter(e *gin.Engine) {
	nodeCtl := iam.NewResourceOperationNodeController()
	node := e.Group("")
	{
		node.GET("/api/quick/resources/:resourceId/operations/:operationId/nodes", nodeCtl.ListResourceOperationNodes)
		node.GET("/api/quick/resources/:resourceId/operations/:operationId/nodes/:nodeId", nodeCtl.GetResourceOperationNode)
		node.POST("/api/quick/resources/:resourceId/operations/:operationId/nodes", nodeCtl.CreateResourceOperationNode)
		node.PUT("/api/quick/resources/:resourceId/operations/:operationId/nodes/:nodeId", nodeCtl.UpdateResourceOperationNode)
		node.DELETE("/api/quick/resources/:resourceId/operations/:operationId/nodes/:nodeId", nodeCtl.DeleteResourceOperationNode)
	}
}

func AddResourceRouter(e *gin.Engine) {
	resourceCtl := iam.NewResourceController()
	resource := e.Group("")
	{
		resource.GET("/api/quick/resources", resourceCtl.ListResources)
		resource.GET("/api/quick/resources/:resourceId", resourceCtl.GetResource)
		resource.POST("/api/quick/resources", resourceCtl.CreateResource)
		resource.PUT("/api/quick/resources/:resourceId", resourceCtl.UpdateResource)
		resource.DELETE("/api/quick/resources/:resourceId", resourceCtl.DeleteResource)
	}
}

func AddResourceNodeRouter(e *gin.Engine) {
	nodeCtl := iam.NewResourceNodeController()
	node := e.Group("")
	{
		node.GET("/api/quick/resources/:resourceId/nodes", nodeCtl.ListResourceNodes)
		node.GET("/api/quick/resources/:resourceId/nodes/:nodeId", nodeCtl.GetResourceNode)
		node.POST("/api/quick/resources/:resourceId/nodes", nodeCtl.CreateResourceNode)
		node.PUT("/api/quick/resources/:resourceId/nodes/:nodeId", nodeCtl.UpdateResourceNode)
		node.DELETE("/api/quick/resources/:resourceId/nodes/:nodeId", nodeCtl.DeleteResourceNode)
	}
}

func AddResourceOperationRouter(e *gin.Engine) {
	operationCtl := iam.NewResourceOperationController()
	operation := e.Group("")
	{
		operation.GET("/api/quick/resources/:resourceId/operations", operationCtl.ListResourceOperations)
		operation.GET("/api/quick/resources/:resourceId/operations/:operationId", operationCtl.GetResourceOperation)
		operation.POST("/api/quick/resources/:resourceId/operations", operationCtl.CreateResourceOperation)
		operation.PUT("/api/quick/resources/:resourceId/operations/:operationId", operationCtl.UpdateResourceOperation)
		operation.DELETE("/api/quick/resources/:resourceId/operations/:operationId", operationCtl.DeleteResourceOperation)
	}
}
