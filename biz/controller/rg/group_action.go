package rg

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/rg"
	"github.com/gin-gonic/gin"
)

// GetResourceGroupActionList
// @Summary	获取资源组动作列表
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	groupId		path	string	true	"group id"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/actions [get]
func GetResourceGroupActionList(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	actions, err := rg.GetResourceGroupActionList(in.Tenant.ID, in.GroupId)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceGroupActionList err")
		return
	}
	resp.SuccessArrayData(c, len(actions), actions)
}

// GetResourceGroupAction
// @Summary	获取资源组动作
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	groupId		path	string	true	"group id"
// @Param	actionId	path	string	true	"action id"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/actions/{actionId} [get]
func GetResourceGroupAction(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	action, err := rg.GetResourceGroupAction(in.Tenant.ID, in.GroupId, in.ActionId)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceGroupAction err")
		return
	}
	resp.SuccessWithData(c, action)
}

// CreateResourceGroupAction
// @Summary	创建资源组动作
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	groupId		path	string	true	"group id"
// @Param	role		body	model.RequestResourceGroup	true	"body"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/actions [post]
func CreateResourceGroupAction(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUriAndJson(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	action, err := rg.CreateResourceGroupAction(in.Tenant.ID, in.GroupId, in.Uid, in.Name, in.Description)
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceGroupAction err")
		return
	}
	resp.SuccessWithData(c, action)
}

// UpdateResourceGroupAction
// @Summary	更新资源组动作
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	groupId		path	string	true	"group id"
// @Param	actionId	path	string	true	"action id"
// @Param	role		body	model.RequestResourceGroup	true	"body"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/actions/{actionId} [put]
func UpdateResourceGroupAction(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUriAndJson(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := rg.UpdateResourceGroupAction(in.Tenant.ID, in.GroupId, in.ActionId, in.Name, in.Description); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceGroupAction err")
		return
	}
	resp.Success(c)
}

// DeleteResourceGroupAction
// @Summary	删除资源组动作
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	groupId		path	string	true	"group id"
// @Param	actionId	path	string	true	"action id"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/actions/{actionId} [delete]
func DeleteResourceGroupAction(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := rg.DeleteResourceGroupAction(in.Tenant.ID, in.GroupId, in.ActionId); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceGroupAction err")
		return
	}
	resp.Success(c)
}
