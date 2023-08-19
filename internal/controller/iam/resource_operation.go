package iam

import (
	"QuickAuth/internal/controller/internal"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/internal/service/iam"

	"github.com/gin-gonic/gin"
)

type ResourceOperation struct {
	internal.Api
}

func NewResourceOperationController() *ResourceOperation {
	return &ResourceOperation{}
}

// ListResourceOperations godoc
// @Summary		list resource operations
// @Tags		resource-operation
// @Param		resourceId	path	string	true	"resource id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/operations 	[get]
func (a ResourceOperation) ListResourceOperations(c *gin.Context) {
	data, err := iam.ListResourceOperations()
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceOperations err", true)
		return
	}

	resp.SuccessArrayData(c, 0, data)
}

// GetResourceOperation godoc
// @Summary		get resource operation
// @Tags		resource-operation
// @Param		resourceId	path	string	true	"resource id"
// @Param		operationId	path	string	true	"operation id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/operations/{operationId} 	[get]
func (a ResourceOperation) GetResourceOperation(c *gin.Context) {
	data, err := iam.GetResourceOperation()
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceOperation err")
		return
	}

	resp.SuccessWithData(c, data)
}

// CreateResourceOperation godoc
// @Summary		create resource operation
// @Tags		resource-operation
// @Param		resourceId	path	string	true	"resource id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/operations 	[post]
func (a ResourceOperation) CreateResourceOperation(c *gin.Context) {
	data, err := iam.CreateResourceOperation()
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceOperation err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResourceOperation godoc
// @Summary		update resource operation
// @Tags		resource-operation
// @Param		resourceId	path	string	true	"resource id"
// @Param		operationId	path	string	true	"operation id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/operations/{operationId} 	[put]
func (a ResourceOperation) UpdateResourceOperation(c *gin.Context) {
	if err := iam.UpdateResourceOperation(); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceOperation err")
		return
	}

	resp.Success(c)
}

// DeleteResourceOperation godoc
// @Summary		delete resource operation
// @Tags		resource-operation
// @Param		resourceId	path	string	true	"resource id"
// @Param		operationId	path	string	true	"operation id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/operations/{operationId} 	[delete]
func (a ResourceOperation) DeleteResourceOperation(c *gin.Context) {
	if err := iam.DeleteResourceOperation(); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceOperation err")
		return
	}

	resp.Success(c)
}
