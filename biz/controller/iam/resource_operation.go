package iam

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/iam"
	"github.com/gin-gonic/gin"
)

// ListResourceOperations
// @Summary	list resource operations
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/operations 	[get]
func ListResourceOperations(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := iam.ListResourceOperations(in.Tenant.ID, in.ResourceId)
	if err != nil {
		resp.ErrorSelect(c, err, "ListResourceOperations err")
		return
	}

	resp.SuccessArrayData(c, len(data), data)
}

// GetResourceOperation
// @Summary	get resource operation
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Param	operationId	path	string	true	"operation id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/operations/{operationId} 	[get]
func GetResourceOperation(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	data, err := iam.GetResourceOperation(in.Tenant.ID, in.ResourceId, in.OperationId)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceOperation err")
		return
	}

	resp.SuccessWithData(c, data)
}

// CreateResourceOperation
// @Summary	create resource operation
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Success	200		{object}	interface{}
// @Router	/api/quick/resources/{resourceId}/operations 	[post]
func CreateResourceOperation(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).BindJson(&in.Operation).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.Operation.TenantID = in.Tenant.ID
	in.Operation.ResourceID = in.ResourceId
	data, err := iam.CreateResourceOperation(&in.Operation)
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceOperation err")
		return
	}

	resp.SuccessWithData(c, data)
}

// UpdateResourceOperation
// @Summary	update resource operation
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Param	operationId	path	string	true	"operation id"
// @Success	200
// @Router	/api/quick/resources/{resourceId}/operations/{operationId} 	[put]
func UpdateResourceOperation(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).BindJson(&in.Operation).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	in.Operation.ID = in.OperationId
	in.Operation.TenantID = in.Tenant.ID
	in.Operation.ResourceID = in.ResourceId
	if err := iam.UpdateResourceOperation(in.Tenant.ID, in.ResourceId, in.OperationId, &in.Operation); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceOperation err")
		return
	}

	resp.Success(c)
}

// DeleteResourceOperation
// @Summary	delete resource operation
// @Tags	ABAC
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	resourceId	path	string	true	"resource id"
// @Param	operationId	path	string	true	"operation id"
// @Success	200
// @Router	/api/quick/resources/{resourceId}/operations/{operationId} 	[delete]
func DeleteResourceOperation(c *gin.Context) {
	var in request.Iam
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := iam.DeleteResourceOperation(in.Tenant.ID, in.ResourceId, in.OperationId); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceOperation err")
		return
	}

	resp.Success(c)
}
