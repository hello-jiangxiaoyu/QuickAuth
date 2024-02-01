package rg

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/pkg/global"
	"gorm.io/gorm"
)

func GetResourceGroupRoleList(tenantId, groupId int64) ([]model.ResourceGroupRole, error) {
	var roles []model.ResourceGroupRole
	if err := global.Db().Where("group_id = ? AND tenant_id = ?", groupId, tenantId).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func GetResourceGroupRole(tenantId, groupId, roleId int64) (*model.ResourceGroupRole, error) {
	var role model.ResourceGroupRole
	if err := global.Db().Where("id = ? AND group_id = ? AND tenant_id = ?", roleId, groupId, tenantId).
		Find(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func CreateResourceGroupRole(tenantId, groupId int64, name string, des string, uid int64) (*model.ResourceGroupRole, error) {
	role := model.ResourceGroupRole{
		Id:          uid,
		TenantId:    tenantId,
		GroupId:     groupId,
		Name:        name,
		Description: des,
	}
	if err := global.Db().Create(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func UpdateResourceGroupRole(tenantId, groupId, roleId int64, name, des string) error {
	if err := global.Db().Model(&model.ResourceGroupRole{}).
		Where("id = ? AND group_id = ? AND tenant_id = ?", roleId, groupId, tenantId).
		Updates(map[string]any{"name": name, "description": des}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteResourceGroupRole(tenantId, groupId, roleId int64) error {
	return global.Db().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("role_id = ? AND tenant_id = ?", roleId, tenantId).
			Delete(&model.ResourceGroupRoleAction{}).Error; err != nil {
			return err
		}
		if err := tx.Where("id = ? AND group_id = ? AND tenant_id = ?", roleId, groupId, tenantId).
			Delete(&model.ResourceGroupRole{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func GetResourceGroupUserList(tenantId, groupId int64) ([]model.ResourceGroupUser, error) {
	var role []model.ResourceGroupUser
	if err := global.Db().Table("resource_group_users as ru").
		Select("ru.id", "ru.group_id", "ru.role_id", "rr.name as role_name", "u.display_name").
		Joins("LEFT JOIN resource_group_roles rr ON rr.id = ru.role_id").
		Joins("LEFT JOIN users u ON u.id = ru.user_id").
		Where("ru.group_id = ? AND ru.tenant_id = ?", groupId, tenantId).Find(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func GetResourceGroupUserRole(tenantId, groupId, userId int64) (*model.ResourceGroupUser, error) {
	var role model.ResourceGroupUser
	if err := global.Db().Where("group_id = ? AND user_id = ? AND tenant_id = ?", groupId, userId, tenantId).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func GetResourceGroupUserActionList(tenantId, groupId, userId int64) ([]model.ResourceGroupRoleAction, error) {
	var userActions []model.ResourceGroupRoleAction
	if err := global.Db().Table("resource_group_role_actions AS rra").
		Select("rra.id", "rra.action_id", "rra.role_id", "rra.tenant_id", "ra.name action_name").
		Joins("LEFT JOIN resource_group_actions AS ra ON ra.id = rra.action_id").
		Joins("LEFT JOIN resource_group_users AS ru ON ru.role_id = rra.role_id").
		Where("ra.group_id = ? AND ru.user_id = ? AND rra.tenant_id", groupId, userId, tenantId).
		Find(&userActions).Error; err != nil {
		return nil, err
	}
	return userActions, nil
}

func GetResourceGroupUserAction(tenantId, userId, actionId int64) (*model.ResourceGroupRoleAction, error) {
	var role model.ResourceGroupRoleAction
	if err := global.Db().Table("resource_group_role_actions as rra").
		Select("rra.id", "rra.action_id", "rra.role_id").
		Joins("LEFT JOIN resource_group_users AS ru ON ru.role_id = rra.role_id").
		Where("rra.action_id = ? AND ru.user_id = ? AND rra.tenant_id = ?", actionId, userId, tenantId).
		First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func CreateResourceGroupUserRole(tenantId, groupId, roleId int64, userIds []int64) error {
	userRoles := make([]model.ResourceGroupUser, 0)
	for _, userId := range userIds {
		userRoles = append(userRoles, model.ResourceGroupUser{
			TenantId: tenantId,
			GroupId:  groupId,
			UserId:   userId,
			RoleId:   roleId,
		})
	}
	if err := global.Db().Create(&userRoles).Error; err != nil {
		return err
	}
	return nil
}

func UpdateResourceGroupUserRole(tenantId, groupId, userId, roleId int64) error {
	if err := global.Db().Model(&model.ResourceGroupUser{}).
		Where("group_id = ? AND user_id = ? AND tenant_id = ?", groupId, userId, tenantId).
		Update("role_id", roleId).Error; err != nil {
		return err
	}
	return nil
}

func DeleteResourceGroupUserRole(tenantId, groupId, userId int64) error {
	if err := global.Db().Where("user_id = ? AND group_id = ? AND tenant_id = ?", userId, groupId, tenantId).
		Delete(&model.ResourceGroupUser{}).Error; err != nil {
		return err
	}
	return nil
}
