package iam

import (
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
	data, err := a.svc.ListResourceOperations()
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceOperations err", true)
		return
	}

	resp.SuccessArrayData(c, 0, data)
}

// GetResourceOperation godoc
// @Summary		get resource operation
// @Tags		resource-node
// @Param		resourceId	path	string	true	"resource id"
// @Param		operationId	path	string	true	"operation id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/operations/{operationId} 	[get]
func (a Resource) GetResourceOperation(c *gin.Context) {
	data, err := a.svc.GetResourceOperation()
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
	data, err := a.svc.CreateResourceOperation()
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
	if err := a.svc.UpdateResourceOperation(); err != nil {
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
	if err := a.svc.DeleteResourceOperation(); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceOperation err")
		return
	}

	resp.Success(c)
}
