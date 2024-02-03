package abac

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/pkg/global"
)

func ListResourceRoleOperations(tenantId int64, resId, roleId int64) ([]model.ResourceRoleOperation, error) {
	var data []model.ResourceRoleOperation
	if err := global.Db().Where("tenant_id = ? AND resource_id = ? AND role_id = ?", tenantId, resId, roleId).
		Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func CreateResourceRoleOperation(roleOption *model.ResourceRoleOperation) (*model.ResourceRoleOperation, error) {
	if err := global.Db().Create(roleOption).Error; err != nil {
		return nil, err
	}
	return roleOption, nil
}

func DeleteResourceRoleOperation(tenantId int64, resId, roleId, optionId int64) error {
	if err := global.Db().Where("id = ? AND tenant_id = ? AND resource_id = ? AND role_id = ?",
		optionId, tenantId, resId, roleId).Delete(model.ResourceRoleOperation{}).Error; err != nil {
		return err
	}

	return nil
}
