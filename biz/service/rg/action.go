package rg

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/pkg/global"
	"gorm.io/gorm"
)

func GetResourceGroupActionList(tenantId, groupId int64) ([]model.ResourceGroupAction, error) {
	var actionsList []model.ResourceGroupAction
	if err := global.Db().Where("group_id = ? AND tenant_id = ?", groupId, tenantId).Find(&actionsList).Error; err != nil {
		return nil, err
	}
	return actionsList, nil
}

func GetResourceGroupAction(tenantId, groupId, actionId int64) (*model.ResourceGroupAction, error) {
	var action model.ResourceGroupAction
	if err := global.Db().Where("id = ? AND group_id = ? AND tenant_id = ?", actionId, groupId, tenantId).First(&action).Error; err != nil {
		return nil, err
	}
	return &action, nil
}

func CreateResourceGroupAction(tenantId int64, groupId int64, uid int64, des string, name string) (*model.ResourceGroupAction, error) {
	action := model.ResourceGroupAction{
		Id:          uid,
		TenantId:    tenantId,
		GroupId:     groupId,
		Name:        name,
		Description: des,
	}
	if err := global.Db().Create(&action).Error; err != nil {
		return nil, err
	}
	return &action, nil
}

func UpdateResourceGroupAction(tenantId, groupId, actionId int64, name, des string) error {
	if err := global.Db().Model(&model.ResourceGroupAction{}).
		Where("id = ? AND group_id = ? AND tenant_id = ?", actionId, groupId, tenantId).
		Updates(map[string]any{"name": name, "description": des}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteResourceGroupAction(tenantId, groupId, actionId int64) error {
	return global.Db().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("action_id = ? AND tenant_id = ?", actionId, tenantId).
			Delete(&model.ResourceGroupRoleAction{}).Error; err != nil {
			return err
		}
		if err := tx.Where("id = ? AND group_id = ? AND tenant_id = ?", actionId, groupId, tenantId).
			Delete(&model.ResourceGroupAction{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func GetResourceGroupRoleActionList(tenantId, roleId int64) ([]model.ResourceGroupRoleAction, error) {
	var actionList []model.ResourceGroupRoleAction
	if err := global.Db().Where("role_id = ? AND tenant_id = ?", roleId, tenantId).Find(&actionList).Error; err != nil {
		return nil, err
	}
	return actionList, nil
}

func GetResourceGroupRoleAction(tenantId, roleId, actionId int64) (*model.ResourceGroupRoleAction, error) {
	var roleAction model.ResourceGroupRoleAction
	if err := global.Db().Where("action_id = ? AND role_id = ? AND tenant_id = ?", actionId, roleId, tenantId).First(&roleAction).Error; err != nil {
		return nil, err
	}
	return &roleAction, nil
}

func CreateResourceGroupRoleAction(tenantId, roleId int64, actionIds []int64) error {
	roleActions := make([]model.ResourceGroupRoleAction, 0)
	for _, actionId := range actionIds {
		roleActions = append(roleActions, model.ResourceGroupRoleAction{
			TenantId: tenantId,
			RoleId:   roleId,
			ActionId: actionId,
		})
	}
	if err := global.Db().Create(&roleActions).Error; err != nil {
		return err
	}
	return nil
}

func DeleteResourceGroupRoleAction(tenantId, roleId int64, actionIds []int64) error {
	if err := global.Db().Where("role_id = ? AND tenant_id = ? AND action_id in ?", roleId, tenantId, actionIds).
		Delete(&model.ResourceGroupRoleAction{}).Error; err != nil {
		return err
	}
	return nil
}
