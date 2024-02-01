package rg

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/rg"
	"github.com/gin-gonic/gin"
)

// GetResourceGroupResourceList
// @Summary	获取资源组的资源列表
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/resources [get]
func GetResourceGroupResourceList(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	res, err := rg.GetResourceGroupResourceList(in.Tenant.ID, in.GroupId)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceGroupResourceList err")
		return
	}
	resp.SuccessArrayData(c, len(res), res)
}

// GetResourceGroupResource
// @Summary	获取资源组的资源
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Param	resourceId	path	string	true	"resource id"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/resources/{resourceId} [get]
func GetResourceGroupResource(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	res, err := rg.GetResourceGroupResource(in.Tenant.ID, in.GroupId, in.ResourceId)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceGroupResource err")
		return
	}
	resp.SuccessWithData(c, res)
}

// CreateResourceGroupResource
// @Summary	创建资源的资源
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Param	data		body	model.RequestResourceGroup	true	"body"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/resources [post]
func CreateResourceGroupResource(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUriAndJson(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	res, err := rg.CreateResourceGroupResource(in.Tenant.ID, in.GroupId, in.Name, in.Description, in.Uid)
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceGroupResource err")
		return
	}
	resp.SuccessWithData(c, res)
}

// UpdateResourceGroupResource
// @Summary	更新资源组的资源
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Param	resourceId	path	string	true	"resource id"
// @Param	data		body	model.RequestResourceGroup	true	"body"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/resources/{resourceId} [put]
func UpdateResourceGroupResource(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUriAndJson(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := rg.UpdateResourceGroupResource(in.Tenant.ID, in.GroupId, in.ResourceId, in.Name, in.Description); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceGroupResource err")
		return
	}
	resp.Success(c)
}

// DeleteResourceGroupResource
// @Summary	删除资源组的资源
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Param	resourceId	path	string	true	"resource id"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/resources/{resourceId} [delete]
func DeleteResourceGroupResource(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := rg.DeleteResourceGroupResource(in.Tenant.ID, in.GroupId, in.ResourceId); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceGroupResource err")
		return
	}
	resp.Success(c)
}
