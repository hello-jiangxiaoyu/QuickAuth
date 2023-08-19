package iam

import (
	"QuickAuth/internal/controller/internal"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/internal/service/iam"

	"github.com/gin-gonic/gin"
)

type ResourceUserRole struct {
	internal.Api
}

func NewResourceUserRoleController() *ResourceUserRole {
	return &ResourceUserRole{}
}

// ListResourceUserRoles godoc
// @Summary		list resource user roles
// @Tags		resource-user-role
// @Param		resourceId	path	string	true	"resource id"
// @Param		userId		path	string	true	"user id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/users/{userId}/roles 	[get]
func (a ResourceUserRole) ListResourceUserRoles(c *gin.Context) {
	data, err := iam.ListResourceUserRoles()
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceUserRoles err", true)
		return
	}

	resp.SuccessArrayData(c, 0, data)
}

// GetResourceUserRole godoc
// @Summary		get resource user role
// @Tags		resource-user-role
// @Param		resourceId	path	string	true	"resource id"
// @Param		userId		path	string	true	"user id"
// @Param		roleId		path	string	true	"role id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/users/{userId}/roles/{roleId} 	[get]
func (a ResourceUserRole) GetResourceUserRole(c *gin.Context) {
	data, err := iam.GetResourceUserRole()
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceUserRole err")
		return
	}

	resp.SuccessWithData(c, data)
}

// CreateResourceUserRole godoc
// @Summary		create resource user role
// @Tags		resource-user-role
// @Param		resourceId	path	string	true	"resource id"
// @Param		userId		path	string	true	"user id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/users/{userId}/roles 	[post]
func (a ResourceUserRole) CreateResourceUserRole(c *gin.Context) {
	data, err := iam.CreateResourceUserRole()
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceUserRole err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResourceUserRole godoc
// @Summary		update resource user role
// @Tags		resource-user-role
// @Param		resourceId	path	string	true	"resource id"
// @Param		userId		path	string	true	"user id"
// @Param		roleId		path	string	true	"role id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/users/{userId}/roles/{roleId} 	[put]
func (a ResourceUserRole) UpdateResourceUserRole(c *gin.Context) {
	if err := iam.UpdateResourceUserRole(); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceUserRole err")
		return
	}

	resp.Success(c)
}

// DeleteResourceUserRole godoc
// @Summary		delete resource user role
// @Tags		resource-user-role
// @Param		resourceId	path	string	true	"resource id"
// @Param		userId		path	string	true	"user id"
// @Param		roleId		path	string	true	"role id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/users/{userId}/roles/{roleId} 	[delete]
func (a ResourceUserRole) DeleteResourceUserRole(c *gin.Context) {
	if err := iam.DeleteResourceUserRole(); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceUserRole err")
		return
	}

	resp.Success(c)
}
