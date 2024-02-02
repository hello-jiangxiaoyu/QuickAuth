package iam

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// ListResourceJSONUserRoles
// @Summary	list resource json user roles
// @Tags	resource-user
// @Param	resourceId	path	string	true	"resource id"
// @Param	userId		path	string	true	"user id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/json/users/{userId}/roles 	[get]
func (a Resource) ListResourceJSONUserRoles(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := a.svc.ListResourceJSONUserRoles(in.Tenant.ID, in.ResourceId, in.UserId)
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceJsonUserRoles err")
		return
	}

	resp.SuccessArrayData(c, len(data), data)
}

// CreateResourceJSONUserRole
// @Summary	create resource json user role
// @Tags	resource-user
// @Param	resourceId	path	string	true	"resource id"
// @Param	userId		path	string	true	"user id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/json/users/{userId}/roles 	[post]
func (a Resource) CreateResourceJSONUserRole(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).BindJson(&in.JsonUserRole).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.JsonUserRole.TenantID = in.Tenant.ID
	in.JsonUserRole.ResourceID = in.ResourceId
	in.JsonUserRole.UserID = in.UserId
	data, err := a.svc.CreateResourceJSONUserRole(&in.JsonUserRole)
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceJsonUserRole err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResourceJSONUserRole
// @Summary	update resource json user role
// @Tags	resource-user
// @Param	resourceId	path	string	true	"resource id"
// @Param	userId		path	string	true	"user id"
// @Param	roleId		path	string	true	"role id"
// @Success	200
// @Router	/api/quick/resources/{resourceId}/json/users/{userId}/roles/{roleId} 	[put]
func (a Resource) UpdateResourceJSONUserRole(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.JsonUserRole.TenantID = in.Tenant.ID
	in.JsonUserRole.ResourceID = in.ResourceId
	in.JsonUserRole.UserID = in.UserId
	if err := a.svc.UpdateResourceJSONUserRole(in.Tenant.ID, in.ResourceId, in.UserId, in.Path, &in.JsonUserRole); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceJsonUserRole err")
		return
	}

	resp.Success(c)
}

// DeleteResourceJSONUserRole
// @Summary	delete resource json user role
// @Tags	resource-user
// @Param	resourceId	path	string	true	"resource id"
// @Param	userId		path	string	true	"user id"
// @Param	roleId		path	string	true	"role id"
// @Success	200
// @Router	/api/quick/resources/{resourceId}/json/users/{userId}/roles/{roleId} 	[delete]
func (a Resource) DeleteResourceJSONUserRole(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := a.svc.DeleteResourceJSONUserRole(in.Tenant.ID, in.ResourceId, in.UserId, in.Path); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceJsonUserRole err")
		return
	}

	resp.Success(c)
}
