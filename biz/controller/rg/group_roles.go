package rg

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/rg"
	"github.com/gin-gonic/gin"
)

// GetResourceGroupRoleList
// @Summary	获取资源组角色列表
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/roles [get]
func GetResourceGroupRoleList(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	res, err := rg.GetResourceGroupRoleList(in.Tenant.ID, in.GroupId)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceGroupRoleList err")
		return
	}
	resp.SuccessArrayData(c, len(res), res)
}

// GetResourceGroupRole
// @Summary	获取资源组角色
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Param	roleId		path	string	true	"role id"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/roles/{roleId} [get]
func GetResourceGroupRole(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	res, err := rg.GetResourceGroupRole(in.Tenant.ID, in.GroupId, in.RoleId)
	if err != nil {
		resp.ErrorSelect(c, err, "GetResourceGroupRole err")
		return
	}
	resp.SuccessWithData(c, res)
}

// CreateResourceGroupRole
// @Summary	创建资源角色
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Param	data		body	model.RequestResourceGroup	true	"body"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/roles [post]
func CreateResourceGroupRole(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUriAndJson(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	res, err := rg.CreateResourceGroupRole(in.Tenant.ID, in.GroupId, in.Name, in.Description, in.Uid)
	if err != nil {
		resp.ErrorCreate(c, err, "CreateResourceGroupRole err")
		return
	}
	resp.SuccessWithData(c, res)
}

// UpdateResourceGroupRole
// @Summary	更新资源组角色
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Param	roleId		path	string	true	"role id"
// @Param	data		body	model.RequestResourceGroup	true	"body"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/roles/{roleId} [put]
func UpdateResourceGroupRole(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUriAndJson(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := rg.UpdateResourceGroupRole(in.Tenant.ID, in.GroupId, in.RoleId, in.Name, in.Description); err != nil {
		resp.ErrorUpdate(c, err, "UpdateResourceGroupRole err")
		return
	}
	resp.Success(c)
}

// DeleteResourceGroupRole
// @Summary	删除资源组角色
// @Tags	resource-group
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	groupId		path	string	true	"group id"
// @Param	roleId		path	string	true	"role id"
// @Success	200
// @Router	/api/quick/resourceGroups/{groupId}/roles/{roleId} [delete]
func DeleteResourceGroupRole(c *gin.Context) {
	var in model.RequestResourceGroup
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := rg.DeleteResourceGroupRole(in.Tenant.ID, in.GroupId, in.RoleId); err != nil {
		resp.ErrorUpdate(c, err, "DeleteResourceGroupRole err")
		return
	}
	resp.Success(c)
}
