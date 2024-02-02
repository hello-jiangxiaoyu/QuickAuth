package iam

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/iam"
	"github.com/gin-gonic/gin"
)

// ListResources
// @Summary	list resources
// @Tags	resource
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources 	[get]
func ListResources(c *gin.Context) {
	var in request.Iam
	if err := internal.New(c).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := iam.ListResources(in.Tenant.ID)
	if err != nil {
		resp.ErrorSelect(c, err, "ListResources err")
		return
	}

	resp.SuccessArrayData(c, len(data), data)
}

// GetResource
// @Summary	get resource
// @Tags	resource
// @Param	resourceId	path	string	true	"resource id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId} 	[get]
func GetResource(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := iam.GetResource(in.Tenant.ID, &in)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResource err")
		return
	}

	resp.SuccessWithData(c, data)
}

// CreateResource
// @Summary	create resource
// @Tags	resource
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources 	[post]
func CreateResource(c *gin.Context) {
	var in request.Iam
	if err := internal.BindJson(c, &in.Resource).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	in.Resource.TenantID = in.Tenant.ID
	data, err := iam.CreateResource(&in.Resource)
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResource err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResource
// @Summary	update resource
// @Tags	resource
// @Param	resourceId	path	string	true	"resource id"
// @Success	200
// @Router	/api/quick/resources/{resourceId} 	[put]
func UpdateResource(c *gin.Context) {
	var in request.Iam
	if err := internal.BindJson(c, &in.Resource).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.Resource.TenantID = in.Tenant.ID
	if err := iam.UpdateResource(in.Tenant.ID, in.ResourceId, &in.Resource); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResource err")
		return
	}

	resp.Success(c)
}

// DeleteResource
// @Summary	delete resource
// @Tags	resource
// @Param	resourceId	path	string	true	"resource id"
// @Success	200
// @Router	/api/quick/resources/{resourceId} 	[delete]
func DeleteResource(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := iam.DeleteResource(in.Tenant.ID, in.ResourceId); err != nil {
		resp.ErrorDelete(c, err, "DeleteResource err")
		return
	}

	resp.Success(c)
}
