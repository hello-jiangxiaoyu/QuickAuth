package abac

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/pkg/global"
)

func ListResourceOperations(tenantId int64, resId int64) ([]model.ResourceOperation, error) {
	var data []model.ResourceOperation
	if err := global.Db().Where("tenant_id = ? AND resource_id = ?", tenantId, resId).
		Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func GetResourceOperation(tenantId int64, resId int64, operationId int64) (*model.ResourceOperation, error) {
	var data model.ResourceOperation
	if err := global.Db().Where("id = ? AND resource_id = ? AND tenant_id = ?", operationId, resId, tenantId).
		First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func CreateResourceOperation(operation *model.ResourceOperation) (*model.ResourceOperation, error) {
	if err := global.Db().Create(operation).Error; err != nil {
		return nil, err
	}

	return operation, nil
}

func UpdateResourceOperation(tenantId int64, resId int64, operationId int64, operation *model.ResourceOperation) error {
	if err := global.Db().Where("id = ? AND resId = ? AND tenant_id = ?", operationId, resId, tenantId).
		Updates(&operation).Error; err != nil {
		return err
	}

	return nil
}

func DeleteResourceOperation(tenantId int64, resId int64, operationId int64) error {
	if err := global.Db().Where("id = ? AND resId = ? AND tenant_id = ?", operationId, resId, tenantId).
		Delete(model.ResourceOperation{}).Error; err != nil {
		return err
	}

	return nil
}
