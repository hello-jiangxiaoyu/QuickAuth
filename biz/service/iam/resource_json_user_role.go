package iam

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/pkg/global"
)

func ListResourceJSONUserRoles(tenantId int64, resId int64, userId string) ([]model.ResourceJSONUserRole, error) {
	var data []model.ResourceJSONUserRole
	if err := global.Db().Where("tenant_id = ? AND resource_id = ? AND user_id = ?", tenantId, resId, userId).
		Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func CreateResourceJSONUserRole(userRole *model.ResourceJSONUserRole) (*model.ResourceJSONUserRole, error) {
	if err := global.Db().Create(userRole).Error; err != nil {
		return nil, err
	}
	return userRole, nil
}

func UpdateResourceJSONUserRole(tenantId int64, resId int64, userId string, path string, userRole *model.ResourceJSONUserRole) error {
	if err := global.Db().Where("tenant_id = ? AND resource_id = ? AND user_id = ? AND json_path = ?",
		tenantId, resId, userId, path).Updates(&userRole).Error; err != nil {
		return err
	}

	return nil
}

func DeleteResourceJSONUserRole(tenantId int64, resId int64, userId string, path string) error {
	if err := global.Db().Where("tenant_id = ? AND resource_id = ? AND user_id = ? AND json_path = ?",
		tenantId, resId, userId, path).Delete(model.ResourceJSONUserRole{}).Error; err != nil {
		return err
	}

	return nil
}
