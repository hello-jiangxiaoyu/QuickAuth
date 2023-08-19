package iam

import (
	"QuickAuth/internal/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// ListResourceRoleOperations godoc
// @Summary		list resource role operations
// @Tags		resource-role-operation
// @Param		resourceId	path	string	true	"resource id"
// @Param		roleId		path	string	true	"role id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/roles/{roleId}/operations 	[get]
func (a Resource) ListResourceRoleOperations(c *gin.Context) {
	data, err := a.svc.ListResourceRoleOperations()
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceRoleOperations err", true)
		return
	}

	resp.SuccessArrayData(c, 0, data)
}

// GetResourceRoleOperation godoc
// @Summary		get resource role operation
// @Tags		resource-role-operation
// @Param		resourceId	path	string	true	"resource id"
// @Param		roleId		path	string	true	"role id"
// @Param		operationId	path	string	true	"operation id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/roles/{roleId}/operations/{operationId} 	[get]
func (a Resource) GetResourceRoleOperation(c *gin.Context) {
	data, err := a.svc.GetResourceRoleOperation()
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceRoleOperation err")
		return
	}

	resp.SuccessWithData(c, data)
}

// CreateResourceRoleOperation godoc
// @Summary		create resource role operation
// @Tags		resource-role-operation
// @Param		resourceId	path	string	true	"resource id"
// @Param		roleId		path	string	true	"role id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/roles/{roleId}/operations 	[post]
func (a Resource) CreateResourceRoleOperation(c *gin.Context) {
	data, err := a.svc.CreateResourceRoleOperation()
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceRoleOperation err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResourceRoleOperation godoc
// @Summary		update resource role operation
// @Tags		resource-role-operation
// @Param		resourceId	path	string	true	"resource id"
// @Param		roleId		path	string	true	"role id"
// @Param		operationId	path	string	true	"operation id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/roles/{roleId}/operations/{operationId} 	[put]
func (a Resource) UpdateResourceRoleOperation(c *gin.Context) {
	if err := a.svc.UpdateResourceRoleOperation(); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceRoleOperation err")
		return
	}

	resp.Success(c)
}

// DeleteResourceRoleOperation godoc
// @Summary		delete resource role operation
// @Tags		resource-role-operation
// @Param		resourceId	path	string	true	"resource id"
// @Param		roleId		path	string	true	"role id"
// @Param		operationId	path	string	true	"operation id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/roles/{roleId}/operations/{operationId} 	[delete]
func (a Resource) DeleteResourceRoleOperation(c *gin.Context) {
	if err := a.svc.DeleteResourceRoleOperation(); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceRoleOperation err")
		return
	}

	resp.Success(c)
}
