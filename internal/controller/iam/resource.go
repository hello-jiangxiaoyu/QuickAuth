package iam

import (
	"QuickAuth/internal/controller/internal"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/internal/service"
	"github.com/gin-gonic/gin"
)

type Resource struct {
	svc *service.Service
	internal.Api
}

func NewResourceController(svc *service.Service) *Resource {
	return &Resource{svc: svc}
}

// ListResources godoc
// @Summary		list resources
// @Tags		resource
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources 	[get]
func (a Resource) ListResources(c *gin.Context) {
	data, err := a.svc.ListResources()
	if err != nil {
		resp.ErrorSelect(c, err, "ListResources err", true)
		return
	}

	resp.SuccessArrayData(c, 0, data)
}

// GetResource godoc
// @Summary		get resource
// @Tags		resource
// @Param		resourceId	path	string	true	"resource id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId} 	[get]
func (a Resource) GetResource(c *gin.Context) {
	data, err := a.svc.GetResource()
	if err != nil {
		resp.ErrorSelect(c, err, "GetResource err")
		return
	}

	resp.SuccessWithData(c, data)
}

// CreateResource godoc
// @Summary		create resource
// @Tags		resource
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources 	[post]
func (a Resource) CreateResource(c *gin.Context) {
	data, err := a.svc.CreateResource()
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResource err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResource godoc
// @Summary		update resource
// @Tags		resource
// @Param		resourceId	path	string	true	"resource id"
// @Success		200
// @Router		/api/quick/resources/{resourceId} 	[put]
func (a Resource) UpdateResource(c *gin.Context) {
	if err := a.svc.UpdateResource(); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResource err")
		return
	}

	resp.Success(c)
}

// DeleteResource godoc
// @Summary		delete resource
// @Tags		resource
// @Param		resourceId	path	string	true	"resource id"
// @Success		200
// @Router		/api/quick/resources/{resourceId} 	[delete]
func (a Resource) DeleteResource(c *gin.Context) {
	if err := a.svc.DeleteResource(); err != nil {
		resp.ErrorDelete(c, err, "DeleteResource err")
		return
	}

	resp.Success(c)
}
