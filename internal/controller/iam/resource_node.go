package iam

import (
	"QuickAuth/internal/endpoint/request"
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
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := a.svc.ListResourceNodes(in.Tenant.ID, in.ResourceId)
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceNodes err", true)
		return
	}

	resp.SuccessArrayData(c, len(data), data)
}

// GetResourceNode godoc
// @Summary		get resource node
// @Tags		resource-node
// @Param		resourceId	path	string	true	"resource id"
// @Param		nodeId		path	string	true	"node id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/nodes/{nodeId} 	[get]
func (a Resource) GetResourceNode(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := a.svc.GetResourceNode(in.Tenant.ID, in.ResourceId, in.NodeId)
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
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).BindJson(&in.Node).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	in.Node.TenantID = in.Tenant.ID
	in.Node.ResourceID = in.ResourceId
	data, err := a.svc.CreateResourceNode(&in.Node)
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
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).BindJson(&in.Node).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.Node.ID = in.NodeId
	in.Node.TenantID = in.Tenant.ID
	in.Node.ResourceID = in.ResourceId
	if err := a.svc.UpdateResourceNode(in.Tenant.ID, in.ResourceId, in.NodeId, &in.Node); err != nil {
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
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := a.svc.DeleteResourceNode(in.Tenant.ID, in.ResourceId, in.NodeId); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceNode err")
		return
	}

	resp.Success(c)
}
