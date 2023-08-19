package iam

import (
	"QuickAuth/internal/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// ListResourceNodes godoc
// @Summary		list resource nodes
// @Tags		resource-node
// @Param		resourceId	path	string	true	"resource id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/nodes 	[get]
func (a Resource) ListResourceNodes(c *gin.Context) {
	data, err := a.svc.ListResourceNodes()
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceNodes err", true)
		return
	}

	resp.SuccessArrayData(c, 0, data)
}

// GetResourceNode godoc
// @Summary		get resource node
// @Tags		resource-node
// @Param		resourceId	path	string	true	"resource id"
// @Param		nodeId		path	string	true	"node id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/nodes/{nodeId} 	[get]
func (a Resource) GetResourceNode(c *gin.Context) {
	data, err := a.svc.GetResourceNode()
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceNode err")
		return
	}

	resp.SuccessWithData(c, data)
}

// CreateResourceNode godoc
// @Summary		create resource node
// @Tags		resource-node
// @Param		resourceId	path	string	true	"resource id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/nodes 	[post]
func (a Resource) CreateResourceNode(c *gin.Context) {
	data, err := a.svc.CreateResourceNode()
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceNode err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResourceNode godoc
// @Summary		update resource node
// @Tags		resource-node
// @Param		resourceId	path	string	true	"resource id"
// @Param		nodeId		path	string	true	"node id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/nodes/{nodeId} 	[put]
func (a Resource) UpdateResourceNode(c *gin.Context) {
	if err := a.svc.UpdateResourceNode(); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceNode err")
		return
	}

	resp.Success(c)
}

// DeleteResourceNode godoc
// @Summary		delete resource node
// @Tags		resource-node
// @Param		resourceId	path	string	true	"resource id"
// @Param		nodeId		path	string	true	"node id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/nodes/{nodeId} 	[delete]
func (a Resource) DeleteResourceNode(c *gin.Context) {
	if err := a.svc.DeleteResourceNode(); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceNode err")
		return
	}

	resp.Success(c)
}
