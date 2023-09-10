package iam

import "QuickAuth/internal/endpoint/model"

func (s *ServiceIam) ListResourceOperations(tenantId int64, resId int64) ([]model.ResourceOperation, error) {
	var data []model.ResourceOperation
	if err := s.db.Where("tenant_id = ? AND resource_id = ?", tenantId, resId).
		Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (s *ServiceIam) GetResourceOperation(tenantId int64, resId int64, operationId int64) (*model.ResourceOperation, error) {
	var data model.ResourceOperation
	if err := s.db.Where("id = ? AND resource_id = ? AND tenant_id = ?", operationId, resId, tenantId).
		First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *ServiceIam) CreateResourceOperation(operation *model.ResourceOperation) (*model.ResourceOperation, error) {
	if err := s.db.Create(operation).Error; err != nil {
		return nil, err
	}

	return operation, nil
}

func (s *ServiceIam) UpdateResourceOperation(tenantId int64, resId int64, operationId int64, operation *model.ResourceOperation) error {
	if err := s.db.Where("id = ? AND resId = ? AND tenant_id = ?", operationId, resId, tenantId).
		Updates(&operation).Error; err != nil {
		return err
	}

	return nil
}

func (s *ServiceIam) DeleteResourceOperation(tenantId int64, resId int64, operationId int64) error {
	if err := s.db.Where("id = ? AND resId = ? AND tenant_id = ?", operationId, resId, tenantId).
		Delete(model.ResourceOperation{}).Error; err != nil {
		return err
	}

	return nil
}
