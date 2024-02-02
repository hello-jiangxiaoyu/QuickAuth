package iam

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/pkg/global"
)

func ListResourceUserRoles(tenantId int64, resId int64, userId string) ([]model.ResourceUserRole, error) {
	var data []model.ResourceUserRole
	if err := global.Db().Where("tenant_id = ? AND resource_id = ? AND user_id = ?", tenantId, resId, userId).
		Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func CreateResourceUserRole(userRole *model.ResourceUserRole) (*model.ResourceUserRole, error) {
	if err := global.Db().Create(userRole).Error; err != nil {
		return nil, err
	}
	return userRole, nil
}

func UpdateResourceUserRole(tenantId int64, resId int64, userId string, nodeId int64, userRole *model.ResourceUserRole) error {
	if err := global.Db().Where("tenant_id = ? AND resource_id = ? AND user_id = ? AND node_id = ?",
		tenantId, resId, userId, nodeId).Updates(&userRole).Error; err != nil {
		return err
	}

	return nil
}

func DeleteResourceUserRole(tenantId int64, resId int64, userId string, nodeId int64) error {
	if err := global.Db().Where("tenant_id = ? AND resource_id = ? AND user_id = ? AND node_id = ?",
		tenantId, resId, userId, nodeId).Delete(model.ResourceUserRole{}).Error; err != nil {
		return err
	}

	return nil
}
