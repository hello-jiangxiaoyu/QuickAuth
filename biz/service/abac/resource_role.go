package abac

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/pkg/global"
)

func ListResourceRoles(tenantId int64, resId int64) ([]model.ResourceRole, error) {
	var data []model.ResourceRole
	if err := global.Db().Where("tenant_id = ? AND resource_id = ?", tenantId, resId).
		Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func GetResourceRole(tenantId int64, resId int64, roleId int64) (*model.ResourceRole, error) {
	var data model.ResourceRole
	if err := global.Db().Where("id = ? AND resource_id = ? AND tenant_id = ?", roleId, resId, tenantId).
		First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func CreateResourceRole(role *model.ResourceRole) (*model.ResourceRole, error) {
	if err := global.Db().Create(role).Error; err != nil {
		return nil, err
	}

	return role, nil
}

func UpdateResourceRole(tenantId int64, resId int64, roleId int64, role *model.ResourceRole) error {
	if err := global.Db().Where("id = ? AND resId = ? AND tenant_id = ?", roleId, resId, tenantId).
		Updates(role).Error; err != nil {
		return err
	}

	return nil
}

func DeleteResourceRole(tenantId int64, resId int64, roleId int64) error {
	if err := global.Db().Where("id = ? AND resId = ? AND tenant_id = ?", roleId, resId, tenantId).
		Delete(model.ResourceRole{}).Error; err != nil {
		return err
	}

	return nil
}
