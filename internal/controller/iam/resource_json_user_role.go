package iam

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// ListResourceJsonUserRoles godoc
// @Summary		list resource json user roles
// @Tags		resource-user
// @Param		resourceId	path	string	true	"resource id"
// @Param		userId		path	string	true	"user id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/json/users/{userId}/roles 	[get]
func (a Resource) ListResourceJsonUserRoles(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := a.svc.ListResourceJsonUserRoles(in.Tenant.ID, in.ResourceId, in.UserId)
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceJsonUserRoles err", true)
		return
	}

	resp.SuccessArrayData(c, len(data), data)
}

// ListResourceOperationNodes godoc
// @Summary		list resource operation nodes
// @Tags		resource-user
// @Param		resourceId	path	string	true	"resource id"
// @Param		operationId	path	string	true	"operation id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/json/operations/{operationId}/nodes 	[get]
func (a Resource) ListResourceOperationNodes(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := a.svc.ListResourceOperationNodes(in.Tenant.ID, in.ResourceId, in.UserId)
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceOperationNodes err", true)
		return
	}

	resp.SuccessArrayData(c, len(data), data)
}

// CreateResourceJsonUserRole godoc
// @Summary		create resource json user role
// @Tags		resource-user
// @Param		resourceId	path	string	true	"resource id"
// @Param		userId		path	string	true	"user id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/json/users/{userId}/roles 	[post]
func (a Resource) CreateResourceJsonUserRole(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).BindJson(&in.JsonUserRole).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.JsonUserRole.TenantID = in.Tenant.ID
	in.JsonUserRole.ResourceID = in.ResourceId
	in.JsonUserRole.UserID = in.UserId
	data, err := a.svc.CreateResourceJsonUserRole(&in.JsonUserRole)
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceJsonUserRole err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResourceJsonUserRole godoc
// @Summary		update resource json user role
// @Tags		resource-user
// @Param		resourceId	path	string	true	"resource id"
// @Param		userId		path	string	true	"user id"
// @Param		roleId		path	string	true	"role id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/json/users/{userId}/roles/{roleId} 	[put]
func (a Resource) UpdateResourceJsonUserRole(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.JsonUserRole.TenantID = in.Tenant.ID
	in.JsonUserRole.ResourceID = in.ResourceId
	in.JsonUserRole.UserID = in.UserId
	if err := a.svc.UpdateResourceJsonUserRole(in.Tenant.ID, in.ResourceId, in.UserId, in.Path, &in.JsonUserRole); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceJsonUserRole err")
		return
	}

	resp.Success(c)
}

// DeleteResourceJsonUserRole godoc
// @Summary		delete resource json user role
// @Tags		resource-user
// @Param		resourceId	path	string	true	"resource id"
// @Param		userId		path	string	true	"user id"
// @Param		roleId		path	string	true	"role id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/json/users/{userId}/roles/{roleId} 	[delete]
func (a Resource) DeleteResourceJsonUserRole(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := a.svc.DeleteResourceJsonUserRole(in.Tenant.ID, in.ResourceId, in.UserId, in.Path); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceJsonUserRole err")
		return
	}

	resp.Success(c)
}
