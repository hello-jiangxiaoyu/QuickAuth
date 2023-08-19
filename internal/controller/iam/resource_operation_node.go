package iam

import (
	"QuickAuth/internal/controller/internal"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/internal/service/iam"

	"github.com/gin-gonic/gin"
)

type ResourceOperationNode struct {
	internal.Api
}

func NewResourceOperationNodeController() *ResourceOperationNode {
	return &ResourceOperationNode{}
}

// ListResourceOperationNodes godoc
// @Summary		list resource operation nodes
// @Tags		resource-operation-node
// @Param		resourceId	path	string	true	"resource id"
// @Param		operationId	path	string	true	"operation id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/operations/{operationId}/nodes 	[get]
func (a ResourceOperationNode) ListResourceOperationNodes(c *gin.Context) {
	data, err := iam.ListResourceOperationNodes()
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceOperationNodes err", true)
		return
	}

	resp.SuccessArrayData(c, 0, data)
}

// GetResourceOperationNode godoc
// @Summary		get resource operation node
// @Tags		resource-operation-node
// @Param		resourceId	path	string	true	"resource id"
// @Param		operationId	path	string	true	"operation id"
// @Param		nodeId		path	string	true	"node id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/operations/{operationId}/nodes/{nodeId} 	[get]
func (a ResourceOperationNode) GetResourceOperationNode(c *gin.Context) {
	data, err := iam.GetResourceOperationNode()
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceOperationNode err")
		return
	}

	resp.SuccessWithData(c, data)
}

// CreateResourceOperationNode godoc
// @Summary		create resource operation node
// @Tags		resource-operation-node
// @Param		resourceId	path	string	true	"resource id"
// @Param		operationId	path	string	true	"operation id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/operations/{operationId}/nodes 	[post]
func (a ResourceOperationNode) CreateResourceOperationNode(c *gin.Context) {
	data, err := iam.CreateResourceOperationNode()
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceOperationNode err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResourceOperationNode godoc
// @Summary		update resource operation node
// @Tags		resource-operation-node
// @Param		resourceId	path	string	true	"resource id"
// @Param		operationId	path	string	true	"operation id"
// @Param		nodeId		path	string	true	"node id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/operations/{operationId}/nodes/{nodeId} 	[put]
func (a ResourceOperationNode) UpdateResourceOperationNode(c *gin.Context) {
	if err := iam.UpdateResourceOperationNode(); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceOperationNode err")
		return
	}

	resp.Success(c)
}

// DeleteResourceOperationNode godoc
// @Summary		delete resource operation node
// @Tags		resource-operation-node
// @Param		resourceId	path	string	true	"resource id"
// @Param		operationId	path	string	true	"operation id"
// @Param		nodeId		path	string	true	"node id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/operations/{operationId}/nodes/{nodeId} 	[delete]
func (a ResourceOperationNode) DeleteResourceOperationNode(c *gin.Context) {
	if err := iam.DeleteResourceOperationNode(); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceOperationNode err")
		return
	}

	resp.Success(c)
}
