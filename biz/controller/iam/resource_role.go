package iam

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/iam"
	"github.com/gin-gonic/gin"
)

// ListResourceRoles
// @Summary	list resource roles
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/roles 	[get]
func ListResourceRoles(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := iam.ListResourceRoles(in.Tenant.ID, in.ResourceId)
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceRoles err")
		return
	}

	resp.SuccessArrayData(c, len(data), data)
}

// GetResourceRole
// @Summary	get resource role
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Param	roleId		path	string	true	"role id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/roles/{roleId} 	[get]
func GetResourceRole(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := iam.GetResourceRole(in.Tenant.ID, in.ResourceId, in.RoleId)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceRole err")
		return
	}

	resp.SuccessWithData(c, data)
}

// CreateResourceRole
// @Summary	create resource role
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/roles 	[post]
func CreateResourceRole(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).BindJson(&in.Role).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.Role.TenantID = in.Tenant.ID
	in.Role.ResourceID = in.ResourceId
	data, err := iam.CreateResourceRole(&in.Role)
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceRole err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResourceRole
// @Summary	update resource role
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Param	roleId		path	string	true	"role id"
// @Success	200
// @Router	/api/quick/resources/{resourceId}/roles/{roleId} 	[put]
func UpdateResourceRole(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).BindJson(&in.Role).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.Role.ID = in.RoleId
	in.Role.TenantID = in.Tenant.ID
	in.Role.ResourceID = in.ResourceId
	if err := iam.UpdateResourceRole(in.Tenant.ID, in.ResourceId, in.RoleId, &in.Role); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceRole err")
		return
	}

	resp.Success(c)
}

// DeleteResourceRole
// @Summary	delete resource role
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Param	roleId		path	string	true	"role id"
// @Success	200
// @Router	/api/quick/resources/{resourceId}/roles/{roleId} 	[delete]
func DeleteResourceRole(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := iam.DeleteResourceRole(in.Tenant.ID, in.ResourceId, in.RoleId); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceRole err")
		return
	}

	resp.Success(c)
}
