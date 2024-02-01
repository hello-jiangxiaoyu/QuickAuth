package iam

import (
	"QuickAuth/biz/endpoint/model"
)

func (s *ServiceIam) ListResourceRoleOperations(tenantId int64, resId, roleId int64) ([]model.ResourceRoleOperation, error) {
	var data []model.ResourceRoleOperation
	if err := s.db.Where("tenant_id = ? AND resource_id = ? AND role_id = ?", tenantId, resId, roleId).
		Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (s *ServiceIam) CreateResourceRoleOperation(roleOption *model.ResourceRoleOperation) (*model.ResourceRoleOperation, error) {
	if err := s.db.Create(roleOption).Error; err != nil {
		return nil, err
	}
	return roleOption, nil
}

func (s *ServiceIam) DeleteResourceRoleOperation(tenantId int64, resId, roleId, optionId int64) error {
	if err := s.db.Where("id = ? AND tenant_id = ? AND resource_id = ? AND role_id = ?",
		optionId, tenantId, resId, roleId).Delete(model.ResourceRoleOperation{}).Error; err != nil {
		return err
	}

	return nil
}
