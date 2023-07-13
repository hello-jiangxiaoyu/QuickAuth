package service

import (
	"QuickAuth/pkg/model"
	"strconv"
)

func (s *Service) GetTenant(appId, tenantId string) (*model.Tenant, error) {
	var tenant model.Tenant
	if err := s.db.Where("id = ? AND app_id = ?", tenantId, appId).
		First(&tenant).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (s *Service) ListTenant(appId string) ([]model.Tenant, error) {
	var tenant []model.Tenant
	if err := s.db.Select("id", "app_id", "user_pool_id", "type", "host", "name", "company", "create_time", "update_time").
		Where("app_id = ?", appId).Find(&tenant).Error; err != nil {
		return nil, err
	}
	return tenant, nil
}

func (s *Service) CreatTenant(t model.Tenant) (*model.Tenant, error) {
	if _, err := s.GetApp(t.AppId); err != nil {
		return nil, err
	}
	if _, err := s.GetUserPool(strconv.FormatInt(t.UserPoolID, 10)); err != nil {
		return nil, err
	}
	if err := s.db.Create(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *Service) ModifyTenant(t model.Tenant) error {
	if err := s.db.Where("id = ? AND app_id = ?", t.ID, t.AppId).Save(&t).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteTenant(appId, tenantId string) error {
	if _, err := s.GetTenant(appId, tenantId); err != nil {
		return err
	}
	if err := s.db.Where("id = ? AND app_id = ?", tenantId, appId).Delete(&model.Tenant{}).Error; err != nil {
		return err
	}
	return nil
}
