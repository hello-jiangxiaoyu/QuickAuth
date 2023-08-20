package iam

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// ListResourceOperations godoc
// @Summary		list resource operations
// @Tags		resource-node
// @Param		resourceId	path	string	true	"resource id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/operations 	[get]
func (a Resource) ListResourceOperations(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := a.svc.ListResourceOperations(in.Tenant.ID, in.ResourceId)
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceOperations err", true)
		return
	}

	resp.SuccessArrayData(c, len(data), data)
}

// GetResourceOperation godoc
// @Summary		get resource operation
// @Tags		resource-node
// @Param		resourceId	path	string	true	"resource id"
// @Param		operationId	path	string	true	"operation id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/operations/{operationId} 	[get]
func (a Resource) GetResourceOperation(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := a.svc.GetResourceOperation(in.Tenant.ID, in.ResourceId, in.OperationId)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceOperation err")
		return
	}

	resp.SuccessWithData(c, data)
}

// CreateResourceOperation godoc
// @Summary		create resource operation
// @Tags		resource-node
// @Param		resourceId	path	string	true	"resource id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/operations 	[post]
func (a Resource) CreateResourceOperation(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).BindJson(&in.Operation).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.Operation.TenantID = in.Tenant.ID
	in.Operation.ResourceID = in.ResourceId
	data, err := a.svc.CreateResourceOperation(&in.Operation)
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceOperation err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResourceOperation godoc
// @Summary		update resource operation
// @Tags		resource-node
// @Param		resourceId	path	string	true	"resource id"
// @Param		operationId	path	string	true	"operation id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/operations/{operationId} 	[put]
func (a Resource) UpdateResourceOperation(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).BindJson(&in.Operation).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.Operation.ID = in.OperationId
	in.Operation.TenantID = in.Tenant.ID
	in.Operation.ResourceID = in.ResourceId
	if err := a.svc.UpdateResourceOperation(in.Tenant.ID, in.ResourceId, in.OperationId, &in.Operation); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceOperation err")
		return
	}

	resp.Success(c)
}

// DeleteResourceOperation godoc
// @Summary		delete resource operation
// @Tags		resource-node
// @Param		resourceId	path	string	true	"resource id"
// @Param		operationId	path	string	true	"operation id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/operations/{operationId} 	[delete]
func (a Resource) DeleteResourceOperation(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := a.svc.DeleteResourceOperation(in.Tenant.ID, in.ResourceId, in.OperationId); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceOperation err")
		return
	}

	resp.Success(c)
}
