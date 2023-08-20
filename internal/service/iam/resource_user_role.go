package iam

import "QuickAuth/internal/endpoint/model"

func (s *ServiceIam) ListResourceUserRoles(tenantId int64, resId int64, userId string) ([]model.ResourceUserRole, error) {
	var data []model.ResourceUserRole
	if err := s.db.Where("tenant_id = ? AND resource_id = ? AND user_id = ?", tenantId, resId, userId).
		Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (s *ServiceIam) GetResourceUserRole(tenantId int64, resId string, userId int64, nodeId int64) (*model.ResourceUserRole, error) {
	var data *model.ResourceUserRole
	if err := s.db.Where("tenant_id = ? AND resource_id = ? AND user_id = ? AND json_path = ?",
		tenantId, resId, userId, nodeId).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (s *ServiceIam) CreateResourceUserRole(userRole *model.ResourceUserRole) (*model.ResourceUserRole, error) {
	if err := s.db.Create(userRole).Error; err != nil {
		return nil, err
	}
	return userRole, nil
}

func (s *ServiceIam) UpdateResourceUserRole(tenantId int64, resId int64, userId string, nodeId int64, userRole *model.ResourceUserRole) error {
	if err := s.db.Where("tenant_id = ? AND resource_id = ? AND user_id = ? AND node_id = ?",
		tenantId, resId, userId, nodeId).Updates(&userRole).Error; err != nil {
		return err
	}

	return nil
}

func (s *ServiceIam) DeleteResourceUserRole(tenantId int64, resId int64, userId string, nodeId int64) error {
	if err := s.db.Where("tenant_id = ? AND resource_id = ? AND user_id = ? AND node_id = ?",
		tenantId, resId, userId, nodeId).Delete(model.ResourceUserRole{}).Error; err != nil {
		return err
	}

	return nil
}
