package rg

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/rg"
	"github.com/gin-gonic/gin"
)

// GetResourceGroupList
// @Summary	获取资源组列表
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Success	200
// @Router	/api/quick/resourceGroups [get]
func GetResourceGroupList(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	groups, err := rg.GetResourceGroupList(in.Tenant.ID)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceGroupList err")
		return
	}
	resp.SuccessArrayData(c, len(groups), groups)
}

// GetResourceGroup
// @Summary	获取资源组详细信息
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	groupId		path	string	true	"group id"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId} [get]
func GetResourceGroup(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	group, err := rg.GetResourceGroup(in.Tenant.ID, in.GroupId)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceGroup err")
		return
	}
	resp.SuccessWithData(c, group)
}

// CreateResourceGroup
// @Summary	创建资源组
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	group		body	model.RequestResourceGroup	true	"body"
// @Success	200
// @Router	/api/quick/resourceGroups [post]
func CreateResourceGroup(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUriAndJson(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	group, err := rg.CreateResourceGroup(in.Tenant.ID, in.Uid, in.Name, in.Description)
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceGroup err")
		return
	}
	resp.SuccessWithData(c, group)
}

// UpdateResourceGroup
// @Summary	更新资源组
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	groupId		path	string	true	"group id"
// @Param	group		body	model.RequestResourceGroup	true	"body"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId} [put]
func UpdateResourceGroup(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUriAndJson(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := rg.UpdateResourceGroup(in.Tenant.ID, in.GroupId, in.Name, in.Description); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceGroup err")
		return
	}
	resp.Success(c)
}

// DeleteResourceGroup
// @Summary	删除资源组
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	groupId		path	string	true	"group id"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId} [delete]
func DeleteResourceGroup(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := rg.DeleteResourceGroup(in.Tenant.ID, in.GroupId); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceGroup err")
		return
	}
	resp.Success(c)
}
