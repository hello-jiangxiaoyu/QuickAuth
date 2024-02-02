package iam

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/iam"
	"github.com/gin-gonic/gin"
)

// ListResourceJSONUserRoles
// @Summary	list resource json user roles
// @Tags	ABAC
// @Param	resourceId	path	string	true	"resource id"
// @Param	userId		path	string	true	"user id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/json/users/{userId}/roles 	[get]
func ListResourceJSONUserRoles(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := iam.ListResourceJSONUserRoles(in.Tenant.ID, in.ResourceId, in.UserId)
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceJsonUserRoles err")
		return
	}

	resp.SuccessArrayData(c, len(data), data)
}

// CreateResourceJSONUserRole
// @Summary	create resource json user role
// @Tags	ABAC
// @Param	resourceId	path	string	true	"resource id"
// @Param	userId		path	string	true	"user id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/json/users/{userId}/roles 	[post]
func CreateResourceJSONUserRole(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).BindJson(&in.JsonUserRole).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.JsonUserRole.TenantID = in.Tenant.ID
	in.JsonUserRole.ResourceID = in.ResourceId
	in.JsonUserRole.UserID = in.UserId
	data, err := iam.CreateResourceJSONUserRole(&in.JsonUserRole)
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceJsonUserRole err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResourceJSONUserRole
// @Summary	update resource json user role
// @Tags	ABAC
// @Param	resourceId	path	string	true	"resource id"
// @Param	userId		path	string	true	"user id"
// @Param	roleId		path	string	true	"role id"
// @Success	200
// @Router	/api/quick/resources/{resourceId}/json/users/{userId}/roles/{roleId} 	[put]
func UpdateResourceJSONUserRole(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.JsonUserRole.TenantID = in.Tenant.ID
	in.JsonUserRole.ResourceID = in.ResourceId
	in.JsonUserRole.UserID = in.UserId
	if err := iam.UpdateResourceJSONUserRole(in.Tenant.ID, in.ResourceId, in.UserId, in.Path, &in.JsonUserRole); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceJsonUserRole err")
		return
	}

	resp.Success(c)
}

// DeleteResourceJSONUserRole
// @Summary	delete resource json user role
// @Tags	ABAC
// @Param	resourceId	path	string	true	"resource id"
// @Param	userId		path	string	true	"user id"
// @Param	roleId		path	string	true	"role id"
// @Success	200
// @Router	/api/quick/resources/{resourceId}/json/users/{userId}/roles/{roleId} 	[delete]
func DeleteResourceJSONUserRole(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := iam.DeleteResourceJSONUserRole(in.Tenant.ID, in.ResourceId, in.UserId, in.Path); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceJsonUserRole err")
		return
	}

	resp.Success(c)
}
