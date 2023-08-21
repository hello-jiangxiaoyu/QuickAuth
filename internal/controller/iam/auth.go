package iam

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// IsOperationAllow godoc
// @Summary		判断用户当前对资源的操作是否被允许
// @Tags		resource-role-operation
// @Param		resourceId	path	string	true	"resource id"
// @Param		nodeId		path	string	true	"node id"
// @Param		operationId	path	string	true	"operation id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/nodes/{nodeId}/operations/{operationId} 	[get]
func (a Resource) IsOperationAllow(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).SetUserInfo().BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	allow, err := a.svc.IsOperationAllow(in.Tenant.ID, in.ResourceId, in.NodeId, in.OperationId, a.UserInfo["sub"])
	if err != nil {
		resp.ErrorSelect(c, err, "IsOperationAllow err")
		return
	}

	resp.SuccessWithData(c, gin.H{"allow": allow})
}

// IsJSONOperationAllow godoc
// @Summary		判断用户当前对JSON资源的操作是否被允许
// @Tags		resource-role-operation
// @Param		resourceId	path	string	true	"resource id"
// @Param		path		query	string	true	"json path"
// @Param		operationId	path	string	true	"operation id"
// @Success		200
// @Router		/api/quick/resources/{resourceId}/json/operations/{operationId} 	[get]
func (a Resource) IsJSONOperationAllow(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).SetUserInfo().BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	allow, err := a.svc.IsJSONOperationAllow(in.Tenant.ID, in.ResourceId, in.Path, in.OperationId, a.UserInfo["sub"])
	if err != nil {
		resp.ErrorSelect(c, err, "IsJSONOperationAllow err")
		return
	}

	resp.SuccessWithData(c, gin.H{"allow": allow})
}

// ListResourceOperationNodes godoc
// @Summary		获取拥有某个操作权限的node列表
// @Tags		resource-user
// @Param		resourceId	path	string	true	"resource id"
// @Param		operationId	path	string	true	"operation id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/json/operations/{operationId}/parents/:parentId 	[get]
func (a Resource) ListResourceOperationNodes(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := a.svc.ListResourceOperationNodes(in.Tenant.ID, in.ResourceId, in.ParentId, in.OperationId, in.UserId)
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceOperationNodes err", true)
		return
	}

	resp.SuccessArrayData(c, len(data), data)
}

// ListJSONResourceOperationNodes godoc
// @Summary		获取拥有某个操作权限的整个json结构
// @Tags		resource-user
// @Param		resourceId	path	string	true	"resource id"
// @Param		operationId	path	string	true	"operation id"
// @Success		200		{object}	interface{}
// @Router		/api/quick/resources/{resourceId}/json/operations/{operationId}/json 	[get]
func (a Resource) ListJSONResourceOperationNodes(c *gin.Context) {
	var in request.Iam
	if err := a.SetCtx(c).SetTenant(&in.Tenant).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := a.svc.ListJSONResourceOperationNodes(in.Tenant.ID, in.ResourceId, in.OperationId, in.UserId)
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceOperationNodes err", true)
		return
	}

	resp.SuccessArrayData(c, len(data), data)
}
