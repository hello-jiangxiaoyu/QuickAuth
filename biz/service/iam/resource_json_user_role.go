package iam

import "QuickAuth/biz/endpoint/model"

func (s *ServiceIam) ListResourceJSONUserRoles(tenantId int64, resId int64, userId string) ([]model.ResourceJSONUserRole, error) {
	var data []model.ResourceJSONUserRole
	if err := s.db.Where("tenant_id = ? AND resource_id = ? AND user_id = ?", tenantId, resId, userId).
		Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (s *ServiceIam) CreateResourceJSONUserRole(userRole *model.ResourceJSONUserRole) (*model.ResourceJSONUserRole, error) {
	if err := s.db.Create(userRole).Error; err != nil {
		return nil, err
	}
	return userRole, nil
}

func (s *ServiceIam) UpdateResourceJSONUserRole(tenantId int64, resId int64, userId string, path string, userRole *model.ResourceJSONUserRole) error {
	if err := s.db.Where("tenant_id = ? AND resource_id = ? AND user_id = ? AND json_path = ?",
		tenantId, resId, userId, path).Updates(&userRole).Error; err != nil {
		return err
	}

	return nil
}

func (s *ServiceIam) DeleteResourceJSONUserRole(tenantId int64, resId int64, userId string, path string) error {
	if err := s.db.Where("tenant_id = ? AND resource_id = ? AND user_id = ? AND json_path = ?",
		tenantId, resId, userId, path).Delete(model.ResourceJSONUserRole{}).Error; err != nil {
		return err
	}

	return nil
}
