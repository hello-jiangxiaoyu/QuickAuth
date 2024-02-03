package abac

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/abac"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// IsOperationAllow
// @Summary	判断用户当前对资源的操作是否被允许
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Param	nodeId		path	string	true	"node id"
// @Param	operationId	path	string	true	"operation id"
// @Success	200
// @Router	/api/quick/resources/{resourceId}/nodes/{nodeId}/operations/{operationId} 	[get]
func IsOperationAllow(c *gin.Context) {
	var in request.Iam
	var user jwt.MapClaims
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).SetUser(&user).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	allow, err := abac.IsOperationAllow(in.Tenant.ID, in.ResourceId, in.NodeId, in.OperationId, user["sub"])
	if err != nil {
		resp.ErrorSelect(c, err, "IsOperationAllow err")
		return
	}

	resp.SuccessWithData(c, gin.H{"allow": allow})
}

// IsJSONOperationAllow
// @Summary	判断用户当前对JSON资源的操作是否被允许
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Param	path		query	string	true	"json path"
// @Param	operationId	path	string	true	"operation id"
// @Success	200
// @Router	/api/quick/resources/{resourceId}/json/operations/{operationId} 	[get]
func IsJSONOperationAllow(c *gin.Context) {
	var in request.Iam
	var user jwt.MapClaims
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).SetUser(&user).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	allow, err := abac.IsJSONOperationAllow(in.Tenant.ID, in.ResourceId, in.Path, in.OperationId, user["sub"])
	if err != nil {
		resp.ErrorSelect(c, err, "IsJSONOperationAllow err")
		return
	}

	resp.SuccessWithData(c, gin.H{"allow": allow})
}

// ListResourceOperationNodes
// @Summary	获取拥有某个操作权限的node列表
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Param	operationId	path	string	true	"operation id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/json/operations/{operationId}/parents/:parentId 	[get]
func ListResourceOperationNodes(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := abac.ListResourceOperationNodes(in.Tenant.ID, in.ResourceId, in.ParentId, in.OperationId, in.UserId)
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceOperationNodes err")
		return
	}

	resp.SuccessArrayData(c, len(data), data)
}

// ListJSONResourceOperationNodes
// @Summary	获取拥有某个操作权限的整个json结构
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Param	operationId	path	string	true	"operation id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/json/operations/{operationId}/json 	[get]
func ListJSONResourceOperationNodes(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := abac.ListJSONResourceOperationNodes(in.Tenant.ID, in.ResourceId, in.OperationId, in.UserId)
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceOperationNodes err")
		return
	}

	resp.SuccessArrayData(c, len(data), data)
}
