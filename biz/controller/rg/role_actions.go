package rg

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/rg"
	"QuickAuth/pkg/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetResourceGroupRoleActionList
// @Summary	获取资源组角色的动作列表
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Param	roleId		path	string	true	"role id"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/roles/{roleId}/actions [get]
func GetResourceGroupRoleActionList(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	res, err := rg.GetResourceGroupRoleActionList(in.Tenant.ID, in.RoleId)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceGroupRoleActionList err")
		return
	}
	resp.SuccessArrayData(c, len(res), res)
}

// GetResourceGroupRoleAction
// @Summary	获取资源组角色的动作
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Param	roleId		path	string	true	"role id"
// @Param	actionId	path	string	true	"action id"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/roles/{roleId}/actions/{actionsId} [get]
func GetResourceGroupRoleAction(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	res, err := rg.GetResourceGroupRoleAction(in.Tenant.ID, in.RoleId, in.ActionId)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceGroupRoleAction err")
		return
	}
	resp.SuccessWithData(c, res)
}

// CreateResourceGroupRoleAction
// @Summary	创建资源角色的动作
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Param	roleId		path	string	true	"role id"
// @Param	data		body	model.RequestResourceGroup	true	"body"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/roles/{roleId}/actions [post]
func CreateResourceGroupRoleAction(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUriAndJson(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if len(in.ActionIds) == 0 {
		resp.ErrorRequestWithMsg(c, "action id should not be empty")
		return
	}
	if err := rg.CreateResourceGroupRoleAction(in.Tenant.ID, in.RoleId, in.ActionIds); err != nil {
		resp.ErrorCreate(c, err, "CreateResourceGroupRoleAction err")
		return
	}
	resp.Success(c)
}

// UpdateResourceGroupRoleAction
// @Summary	更新资源组角色的动作
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Param	roleId		path	string	true	"role id"
// @Param	data		body	model.RequestResourceGroup	true	"body"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/roles/{roleId}/actions [put]
func UpdateResourceGroupRoleAction(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUriAndJson(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if len(in.ActionIds) == 0 {
		resp.ErrorRequestWithMsg(c, "action id should not be empty")
		return
	}

	roleActions := make([]model.ResourceGroupRoleAction, 0)
	for _, actionId := range in.ActionIds {
		roleActions = append(roleActions, model.ResourceGroupRoleAction{
			TenantId: in.Tenant.ID,
			RoleId:   in.RoleId,
			ActionId: actionId,
		})
	}
	if err := global.DB.Transaction(func(tx *gorm.DB) error {
		if err := global.DB.Where("role_id = ? AND tenant_id = ?", in.RoleId, in.Tenant.ID).
			Delete(&model.ResourceGroupRoleAction{}).Error; err != nil {
			return err
		}
		if err := tx.Create(&roleActions).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceGroupRoleAction err")
		return
	}

	resp.Success(c)
}

// DeleteResourceGroupRoleAction
// @Summary	删除资源组角色的动作
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Param	roleId		path	string	true	"role id"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/roles/{roleId}/actions [delete]
func DeleteResourceGroupRoleAction(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if len(in.ActionIds) == 0 {
		resp.ErrorRequestWithMsg(c, "action id should not be empty")
		return
	}
	if err := rg.DeleteResourceGroupRoleAction(in.Tenant.ID, in.RoleId, in.ActionIds); err != nil {
		resp.ErrorCreate(c, err, "DeleteResourceGroupRoleAction err")
		return
	}
	resp.Success(c)
}
