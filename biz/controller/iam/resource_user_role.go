package iam

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/iam"
	"github.com/gin-gonic/gin"
)

// ListResourceUserRoles
// @Summary	list resource json user roles
// @Tags	ABAC
// @Param	resourceId	path	string	true	"resource id"
// @Param	userId		path	string	true	"user id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/users/{userId}/roles 	[get]
func ListResourceUserRoles(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := iam.ListResourceUserRoles(in.Tenant.ID, in.ResourceId, in.UserId)
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceJsonUserRoles err")
		return
	}

	resp.SuccessArrayData(c, len(data), data)
}

// CreateResourceUserRole
// @Summary	create resource json user role
// @Tags	ABAC
// @Param	resourceId	path	string	true	"resource id"
// @Param	userId		path	string	true	"user id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/users/{userId}/roles 	[post]
func CreateResourceUserRole(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).BindJson(&in.JsonUserRole).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.JsonUserRole.TenantID = in.Tenant.ID
	in.JsonUserRole.ResourceID = in.ResourceId
	in.JsonUserRole.UserID = in.UserId
	data, err := iam.CreateResourceUserRole(&in.UserRole)
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceJsonUserRole err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResourceUserRole
// @Summary	update resource json user role
// @Tags	ABAC
// @Param	resourceId	path	string	true	"resource id"
// @Param	userId		path	string	true	"user id"
// @Param	roleId		path	string	true	"role id"
// @Success	200
// @Router	/api/quick/resources/{resourceId}/users/{userId}/roles/{roleId} 	[put]
func UpdateResourceUserRole(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.JsonUserRole.TenantID = in.Tenant.ID
	in.JsonUserRole.ResourceID = in.ResourceId
	in.JsonUserRole.UserID = in.UserId
	if err := iam.UpdateResourceUserRole(in.Tenant.ID, in.ResourceId, in.UserId, in.NodeId, &in.UserRole); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceJsonUserRole err")
		return
	}

	resp.Success(c)
}

// DeleteResourceUserRole
// @Summary	delete resource json user role
// @Tags	ABAC
// @Param	resourceId	path	string	true	"resource id"
// @Param	userId		path	string	true	"user id"
// @Param	roleId		path	string	true	"role id"
// @Success	200
// @Router	/api/quick/resources/{resourceId}/users/{userId}/roles/{roleId} 	[delete]
func DeleteResourceUserRole(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := iam.DeleteResourceUserRole(in.Tenant.ID, in.ResourceId, in.UserId, in.NodeId); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceJsonUserRole err")
		return
	}

	resp.Success(c)
}
