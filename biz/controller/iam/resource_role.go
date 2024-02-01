package iam

import (
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// ListResourceRoles
// @Summary	list resource roles
// @Tags	resource-role
// @Param	resourceId	path	string	true	"resource id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/roles 	[get]
func (a Resource) ListResourceRoles(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := a.svc.ListResourceRoles(in.Tenant.ID, in.ResourceId)
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceRoles err", true)
		return
	}

	resp.SuccessArrayData(c, len(data), data)
}

// GetResourceRole
// @Summary	get resource role
// @Tags	resource-role
// @Param	resourceId	path	string	true	"resource id"
// @Param	roleId		path	string	true	"role id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/roles/{roleId} 	[get]
func (a Resource) GetResourceRole(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := a.svc.GetResourceRole(in.Tenant.ID, in.ResourceId, in.RoleId)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceRole err")
		return
	}

	resp.SuccessWithData(c, data)
}

// CreateResourceRole
// @Summary	create resource role
// @Tags	resource-role
// @Param	resourceId	path	string	true	"resource id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/roles 	[post]
func (a Resource) CreateResourceRole(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).BindJson(&in.Role).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.Role.TenantID = in.Tenant.ID
	in.Role.ResourceID = in.ResourceId
	data, err := a.svc.CreateResourceRole(&in.Role)
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceRole err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResourceRole
// @Summary	update resource role
// @Tags	resource-role
// @Param	resourceId	path	string	true	"resource id"
// @Param	roleId		path	string	true	"role id"
// @Success	200
// @Router	/api/quick/resources/{resourceId}/roles/{roleId} 	[put]
func (a Resource) UpdateResourceRole(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).BindJson(&in.Role).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.Role.ID = in.RoleId
	in.Role.TenantID = in.Tenant.ID
	in.Role.ResourceID = in.ResourceId
	if err := a.svc.UpdateResourceRole(in.Tenant.ID, in.ResourceId, in.RoleId, &in.Role); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceRole err")
		return
	}

	resp.Success(c)
}

// DeleteResourceRole
// @Summary	delete resource role
// @Tags	resource-role
// @Param	resourceId	path	string	true	"resource id"
// @Param	roleId		path	string	true	"role id"
// @Success	200
// @Router	/api/quick/resources/{resourceId}/roles/{roleId} 	[delete]
func (a Resource) DeleteResourceRole(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := a.svc.DeleteResourceRole(in.Tenant.ID, in.ResourceId, in.RoleId); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceRole err")
		return
	}

	resp.Success(c)
}
