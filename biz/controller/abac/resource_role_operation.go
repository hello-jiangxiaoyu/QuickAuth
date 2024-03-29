package abac

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/abac"
	"github.com/gin-gonic/gin"
)

// ListResourceRoleOperations
// @Summary	list resource role operations
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Param	roleId		path	string	true	"role id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/roles/{roleId}/operations 	[get]
func ListResourceRoleOperations(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := abac.ListResourceRoleOperations(in.Tenant.ID, in.ResourceId, in.RoleId)
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceRoleOperations err")
		return
	}

	resp.SuccessArrayData(c, len(data), data)
}

// CreateResourceRoleOperation
// @Summary	create resource role operation
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Param	roleId		path	string	true	"role id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/roles/{roleId}/operations 	[post]
func CreateResourceRoleOperation(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).BindJson(&in.RoleOperation).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.RoleOperation.TenantID = in.Tenant.ID
	in.RoleOperation.ResourceID = in.ResourceId
	in.RoleOperation.RoleID = in.RoleId
	data, err := abac.CreateResourceRoleOperation(&in.RoleOperation)
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceRoleOperation err")
		return
	}

	resp.SuccessWithData(c, data)
}

// DeleteResourceRoleOperation
// @Summary	delete resource role operation
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Param	roleId		path	string	true	"role id"
// @Param	operationId	path	string	true	"operation id"
// @Success	200
// @Router	/api/quick/resources/{resourceId}/roles/{roleId}/operations/{operationId} 	[delete]
func DeleteResourceRoleOperation(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := abac.DeleteResourceRoleOperation(in.Tenant.ID, in.ResourceId, in.RoleId, in.OperationId); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceRoleOperation err")
		return
	}

	resp.Success(c)
}
