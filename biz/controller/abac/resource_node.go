package abac

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/abac"
	"github.com/gin-gonic/gin"
)

// ListResourceNodes
// @Summary	list resource nodes
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/nodes 	[get]
func ListResourceNodes(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := abac.ListResourceNodes(in.Tenant.ID, in.ResourceId)
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceNodes err")
		return
	}

	resp.SuccessArrayData(c, len(data), data)
}

// GetResourceNode
// @Summary	get resource node
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Param	nodeId		path	string	true	"node id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/nodes/{nodeId} 	[get]
func GetResourceNode(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := abac.GetResourceNode(in.Tenant.ID, in.ResourceId, in.NodeId)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceNode err")
		return
	}

	resp.SuccessWithData(c, data)
}

// CreateResourceNode
// @Summary	create resource node
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/nodes 	[post]
func CreateResourceNode(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).BindJson(&in.Node).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	in.Node.TenantID = in.Tenant.ID
	in.Node.ResourceID = in.ResourceId
	data, err := abac.CreateResourceNode(&in.Node)
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceNode err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResourceNode
// @Summary	update resource node
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Param	nodeId		path	string	true	"node id"
// @Success	200
// @Router	/api/quick/resources/{resourceId}/nodes/{nodeId} 	[put]
func UpdateResourceNode(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).BindJson(&in.Node).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.Node.ID = in.NodeId
	in.Node.TenantID = in.Tenant.ID
	in.Node.ResourceID = in.ResourceId
	if err := abac.UpdateResourceNode(in.Tenant.ID, in.ResourceId, in.NodeId, &in.Node); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceNode err")
		return
	}

	resp.Success(c)
}

// DeleteResourceNode
// @Summary	delete resource node
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Param	nodeId		path	string	true	"node id"
// @Success	200
// @Router	/api/quick/resources/{resourceId}/nodes/{nodeId} 	[delete]
func DeleteResourceNode(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := abac.DeleteResourceNode(in.Tenant.ID, in.ResourceId, in.NodeId); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceNode err")
		return
	}

	resp.Success(c)
}
