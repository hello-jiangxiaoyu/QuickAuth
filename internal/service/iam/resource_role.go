package iam

import "QuickAuth/internal/endpoint/model"

func (s *ServiceIam) ListResourceRoles(tenantId int64, resId int64) ([]model.ResourceRole, error) {
	var data []model.ResourceRole
	if err := s.db.Where("tenant_id = ? AND resource_id = ?", tenantId, resId).
		Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (s *ServiceIam) GetResourceRole(tenantId int64, resId int64, roleId int64) (*model.ResourceRole, error) {
	var data model.ResourceRole
	if err := s.db.Where("id = ? AND resource_id = ? AND tenant_id = ?", roleId, resId, tenantId).
		First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *ServiceIam) CreateResourceRole(role *model.ResourceRole) (*model.ResourceRole, error) {
	if err := s.db.Create(role).Error; err != nil {
		return nil, err
	}

	return role, nil
}

func (s *ServiceIam) UpdateResourceRole(tenantId int64, resId int64, roleId int64, role *model.ResourceRole) error {
	if err := s.db.Where("id = ? AND resId = ? AND tenant_id = ?", roleId, resId, tenantId).
		Updates(role).Error; err != nil {
		return err
	}

	return nil
}

func (s *ServiceIam) DeleteResourceRole(tenantId int64, resId int64, roleId int64) error {
	if err := s.db.Where("id = ? AND resId = ? AND tenant_id = ?", roleId, resId, tenantId).
		Delete(model.ResourceRole{}).Error; err != nil {
		return err
	}

	return nil
}
