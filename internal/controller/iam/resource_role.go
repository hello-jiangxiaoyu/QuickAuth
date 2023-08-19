package iam

import (
	"QuickAuth/internal/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// ListResourceRoles godoc
// @Summary		list resource roles
// @Tags		resource-node
// @Param		resourceId	path	string	true	"resource id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/roles 	[get]
func (a Resource) ListResourceRoles(c *gin.Context) {
	data, err := a.svc.ListResourceRoles()
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceRoles err", true)
		return
	}

	resp.SuccessArrayData(c, 0, data)
}

// GetResourceRole godoc
// @Summary		get resource role
// @Tags		resource-node
// @Param		resourceId	path	string	true	"resource id"
// @Param		roleId		path	string	true	"role id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/roles/{roleId} 	[get]
func (a Resource) GetResourceRole(c *gin.Context) {
	data, err := a.svc.GetResourceRole()
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceRole err")
		return
	}

	resp.SuccessWithData(c, data)
}

// CreateResourceRole godoc
// @Summary		create resource role
// @Tags		resource-node
// @Param		resourceId	path	string	true	"resource id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/roles 	[post]
func (a Resource) CreateResourceRole(c *gin.Context) {
	data, err := a.svc.CreateResourceRole()
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceRole err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResourceRole godoc
// @Summary		update resource role
// @Tags		resource-node
// @Param		resourceId	path	string	true	"resource id"
// @Param		roleId		path	string	true	"role id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/roles/{roleId} 	[put]
func (a Resource) UpdateResourceRole(c *gin.Context) {
	if err := a.svc.UpdateResourceRole(); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceRole err")
		return
	}

	resp.Success(c)
}

// DeleteResourceRole godoc
// @Summary		delete resource role
// @Tags		resource-node
// @Param		resourceId	path	string	true	"resource id"
// @Param		roleId		path	string	true	"role id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/roles/{roleId} 	[delete]
func (a Resource) DeleteResourceRole(c *gin.Context) {
	if err := a.svc.DeleteResourceRole(); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceRole err")
		return
	}

	resp.Success(c)
}
