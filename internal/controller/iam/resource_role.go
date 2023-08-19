package iam

import (
	"QuickAuth/internal/controller/internal"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/internal/service/iam"

	"github.com/gin-gonic/gin"
)

type ResourceRole struct {
	internal.Api
}

func NewResourceRoleController() *ResourceRole {
	return &ResourceRole{}
}

// ListResourceRoles godoc
// @Summary		list resource roles
// @Tags		resource-role
// @Param		resourceId	path	string	true	"resource id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/roles 	[get]
func (a ResourceRole) ListResourceRoles(c *gin.Context) {
	data, err := iam.ListResourceRoles()
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceRoles err", true)
		return
	}

	resp.SuccessArrayData(c, 0, data)
}

// GetResourceRole godoc
// @Summary		get resource role
// @Tags		resource-role
// @Param		resourceId	path	string	true	"resource id"
// @Param		roleId		path	string	true	"role id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/roles/{roleId} 	[get]
func (a ResourceRole) GetResourceRole(c *gin.Context) {
	data, err := iam.GetResourceRole()
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceRole err")
		return
	}

	resp.SuccessWithData(c, data)
}

// CreateResourceRole godoc
// @Summary		create resource role
// @Tags		resource-role
// @Param		resourceId	path	string	true	"resource id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/roles 	[post]
func (a ResourceRole) CreateResourceRole(c *gin.Context) {
	data, err := iam.CreateResourceRole()
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceRole err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResourceRole godoc
// @Summary		update resource role
// @Tags		resource-role
// @Param		resourceId	path	string	true	"resource id"
// @Param		roleId		path	string	true	"role id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/roles/{roleId} 	[put]
func (a ResourceRole) UpdateResourceRole(c *gin.Context) {
	if err := iam.UpdateResourceRole(); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceRole err")
		return
	}

	resp.Success(c)
}

// DeleteResourceRole godoc
// @Summary		delete resource role
// @Tags		resource-role
// @Param		resourceId	path	string	true	"resource id"
// @Param		roleId		path	string	true	"role id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/roles/{roleId} 	[delete]
func (a ResourceRole) DeleteResourceRole(c *gin.Context) {
	if err := iam.DeleteResourceRole(); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceRole err")
		return
	}

	resp.Success(c)
}
