package rg

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/rg"
	"errors"
	"github.com/gin-gonic/gin"
)

// GetResourceGroupUserList
// @Summary	组内用户列表
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/users [get]
func GetResourceGroupUserList(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	res, err := rg.GetResourceGroupUserList(in.Tenant.ID, in.GroupId)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceGroupUserList err")
		return
	}
	resp.SuccessArrayData(c, len(res), res)
}

// GetResourceGroupUserRole
// @Summary	用户在组内的角色
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Param	userId		path	integer	true	"client user id"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/users/{userId}/roles [get]
func GetResourceGroupUserRole(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	res, err := rg.GetResourceGroupUserRole(in.Tenant.ID, in.GroupId, in.UserId)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceGroupUserRole err")
		return
	}
	resp.SuccessWithData(c, res)
}

// GetResourceGroupUserActionList
// @Summary	用户在组内所拥有的权限列表
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Param	userId		path	integer	true	"client user id"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/users/{userId}/actions [get]
func GetResourceGroupUserActionList(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	res, err := rg.GetResourceGroupUserActionList(in.Tenant.ID, in.GroupId, in.UserId)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceGroupUserActionList err")
		return
	}
	resp.SuccessArrayData(c, len(res), res)
}

// GetResourceGroupUserAction
// @Summary	用户在组内是否拥有某个权限
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Param	userId		path	string	true	"client user id"
// @Param	actionId	path	string	true	"action id"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/users/{userId}/actions/{actionId} [get]
func GetResourceGroupUserAction(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	res, err := rg.GetResourceGroupUserAction(in.Tenant.ID, in.UserId, in.ActionId)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceGroupUserAction err")
		return
	}
	resp.SuccessWithData(c, res)
}

// CreateResourceGroupUserRole
// @Summary	将用户拉入组内
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Param	roleId		path	string	true	"role id"
// @Param	group		body	model.RequestResourceGroup	true	"body"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/roles/{roleId}/users [post]
func CreateResourceGroupUserRole(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUriAndJson(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if len(in.UserIds) == 0 {
		resp.ErrorRequestWithMsg(c, "user id should not be empty")
		return
	}
	if err := rg.CreateResourceGroupUserRole(in.Tenant.ID, in.GroupId, in.RoleId, in.UserIds); err != nil {
		resp.ErrorCreate(c, err, "CreateResourceGroupUserRole err")
		return
	}
	resp.Success(c)
}

// UpdateResourceGroupUserRole
// @Summary	修改用户在组内的角色
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Param	userId		path	integer	true	"client user id"
// @Param	group		body	model.RequestResourceGroup	true	"body"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/users/{userId} [put]
func UpdateResourceGroupUserRole(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUriAndJson(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if in.RoleId == 0 {
		resp.ErrorRequest(c, errors.New("body roleId should not be empty"))
		return
	}
	if err := rg.UpdateResourceGroupUserRole(in.Tenant.ID, in.GroupId, in.UserId, in.RoleId); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceGroupUserRole err")
		return
	}
	resp.Success(c)
}

// DeleteResourceGroupUser
// @Summary	踢出用户
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Param	userId		path	integer	true	"client user id"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/users/{userId} [delete]
func DeleteResourceGroupUser(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := rg.DeleteResourceGroupUserRole(in.Tenant.ID, in.GroupId, in.UserId); err != nil {
		resp.ErrorDelete(c, err, "DeleteResourceGroupUserRole err")
		return
	}
	resp.Success(c)
}
