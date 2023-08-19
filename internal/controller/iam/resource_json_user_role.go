package iam

import (
	"QuickAuth/internal/controller/internal"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/internal/service/iam"

	"github.com/gin-gonic/gin"
)

type ResourceJsonUserRole struct {
	internal.Api
}

func NewResourceJsonUserRoleController() *ResourceJsonUserRole {
	return &ResourceJsonUserRole{}
}

// ListResourceJsonUserRoles godoc
// @Summary		list resource json user roles
// @Tags		resource-json-user-role
// @Param		resourceId	path	string	true	"resource id"
// @Param		userId		path	string	true	"user id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/json/users/{userId}/roles 	[get]
func (a ResourceJsonUserRole) ListResourceJsonUserRoles(c *gin.Context) {
	data, err := iam.ListResourceJsonUserRoles()
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceJsonUserRoles err", true)
		return
	}

	resp.SuccessArrayData(c, 0, data)
}

// GetResourceJsonUserRole godoc
// @Summary		get resource json user role
// @Tags		resource-json-user-role
// @Param		resourceId	path	string	true	"resource id"
// @Param		userId		path	string	true	"user id"
// @Param		roleId		path	string	true	"role id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/json/users/{userId}/roles/{roleId} 	[get]
func (a ResourceJsonUserRole) GetResourceJsonUserRole(c *gin.Context) {
	data, err := iam.GetResourceJsonUserRole()
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceJsonUserRole err")
		return
	}

	resp.SuccessWithData(c, data)
}

// CreateResourceJsonUserRole godoc
// @Summary		create resource json user role
// @Tags		resource-json-user-role
// @Param		resourceId	path	string	true	"resource id"
// @Param		userId		path	string	true	"user id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/json/users/{userId}/roles 	[post]
func (a ResourceJsonUserRole) CreateResourceJsonUserRole(c *gin.Context) {
	data, err := iam.CreateResourceJsonUserRole()
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceJsonUserRole err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResourceJsonUserRole godoc
// @Summary		update resource json user role
// @Tags		resource-json-user-role
// @Param		resourceId	path	string	true	"resource id"
// @Param		userId		path	string	true	"user id"
// @Param		roleId		path	string	true	"role id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/json/users/{userId}/roles/{roleId} 	[put]
func (a ResourceJsonUserRole) UpdateResourceJsonUserRole(c *gin.Context) {
	if err := iam.UpdateResourceJsonUserRole(); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceJsonUserRole err")
		return
	}

	resp.Success(c)
}

// DeleteResourceJsonUserRole godoc
// @Summary		delete resource json user role
// @Tags		resource-json-user-role
// @Param		resourceId	path	string	true	"resource id"
// @Param		userId		path	string	true	"user id"
// @Param		roleId		path	string	true	"role id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/json/users/{userId}/roles/{roleId} 	[delete]
func (a ResourceJsonUserRole) DeleteResourceJsonUserRole(c *gin.Context) {
	if err := iam.DeleteResourceJsonUserRole(); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceJsonUserRole err")
		return
	}

	resp.Success(c)
}
