package iam

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service"
	"github.com/gin-gonic/gin"
)

type Resource struct {
	svc *service.Service
}

func NewResourceController(svc *service.Service) *Resource {
	return &Resource{svc: svc}
}

// ListResources
// @Summary	list resources
// @Tags	resource
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources 	[get]
func (a Resource) ListResources(c *gin.Context) {
	var in request.Iam
	if err := internal.New(c).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := a.svc.ListResources(in.Tenant.ID)
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
func (a Resource) GetResource(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := a.svc.GetResource(in.Tenant.ID, &in)
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
func (a Resource) CreateResource(c *gin.Context) {
	var in request.Iam
	if err := internal.BindJson(c, &in.Resource).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	in.Resource.TenantID = in.Tenant.ID
	data, err := a.svc.CreateResource(&in.Resource)
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
func (a Resource) UpdateResource(c *gin.Context) {
	var in request.Iam
	if err := internal.BindJson(c, &in.Resource).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.Resource.TenantID = in.Tenant.ID
	if err := a.svc.UpdateResource(in.Tenant.ID, in.ResourceId, &in.Resource); err != nil {
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
func (a Resource) DeleteResource(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := a.svc.DeleteResource(in.Tenant.ID, in.ResourceId); err != nil {
		resp.ErrorDelete(c, err, "DeleteResource err")
		return
	}

	resp.Success(c)
}
