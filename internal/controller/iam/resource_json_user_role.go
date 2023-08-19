package iam

import (
	"QuickAuth/internal/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// ListResourceJsonUserRoles godoc
// @Summary		list resource json user roles
// @Tags		resource-user
// @Param		resourceId	path	string	true	"resource id"
// @Param		userId		path	string	true	"user id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/users/{userId}/roles 	[get]
func (a Resource) ListResourceJsonUserRoles(c *gin.Context) {
	data, err := a.svc.ListResourceJsonUserRoles()
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceJsonUserRoles err", true)
		return
	}

	resp.SuccessArrayData(c, 0, data)
}

// ListResourceOperationNodes godoc
// @Summary		list resource operation nodes
// @Tags		resource-user
// @Param		resourceId	path	string	true	"resource id"
// @Param		operationId	path	string	true	"operation id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/operations/{operationId}/nodes 	[get]
func (a Resource) ListResourceOperationNodes(c *gin.Context) {
	data, err := a.svc.ListResourceOperationNodes()
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceOperationNodes err", true)
		return
	}

	resp.SuccessArrayData(c, 0, data)
}

// GetResourceJsonUserRole godoc
// @Summary		get resource json user role
// @Tags		resource-user
// @Param		resourceId	path	string	true	"resource id"
// @Param		userId		path	string	true	"user id"
// @Param		roleId		path	string	true	"role id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/users/{userId}/roles/{roleId} 	[get]
func (a Resource) GetResourceJsonUserRole(c *gin.Context) {
	data, err := a.svc.GetResourceJsonUserRole()
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceJsonUserRole err")
		return
	}

	resp.SuccessWithData(c, data)
}

// CreateResourceJsonUserRole godoc
// @Summary		create resource json user role
// @Tags		resource-user
// @Param		resourceId	path	string	true	"resource id"
// @Param		userId		path	string	true	"user id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/users/{userId}/roles 	[post]
func (a Resource) CreateResourceJsonUserRole(c *gin.Context) {
	data, err := a.svc.CreateResourceJsonUserRole()
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceJsonUserRole err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResourceJsonUserRole godoc
// @Summary		update resource json user role
// @Tags		resource-user
// @Param		resourceId	path	string	true	"resource id"
// @Param		userId		path	string	true	"user id"
// @Param		roleId		path	string	true	"role id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/users/{userId}/roles/{roleId} 	[put]
func (a Resource) UpdateResourceJsonUserRole(c *gin.Context) {
	if err := a.svc.UpdateResourceJsonUserRole(); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceJsonUserRole err")
		return
	}

	resp.Success(c)
}

// DeleteResourceJsonUserRole godoc
// @Summary		delete resource json user role
// @Tags		resource-user
// @Param		resourceId	path	string	true	"resource id"
// @Param		userId		path	string	true	"user id"
// @Param		roleId		path	string	true	"role id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/users/{userId}/roles/{roleId} 	[delete]
func (a Resource) DeleteResourceJsonUserRole(c *gin.Context) {
	if err := a.svc.DeleteResourceJsonUserRole(); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceJsonUserRole err")
		return
	}

	resp.Success(c)
}
