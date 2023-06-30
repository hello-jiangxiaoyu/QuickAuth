package service

import (
	"QuickAuth/pkg/model"
	"strconv"
)

func (s *Service) GetTenant(clientId, tenantId string) (*model.Tenant, error) {
	var tenant model.Tenant
	if err := s.db.Where("id = ? AND client_id = ?", tenantId, clientId).
		First(&tenant).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (s *Service) ListTenant(clientId string) ([]model.Tenant, error) {
	var tenant []model.Tenant
	if err := s.db.Select("id", "client_id", "user_pool_id", "type", "host", "name", "company", "create_time", "update_time").
		Where("client_id = ?", clientId).Find(&tenant).Error; err != nil {
		return nil, err
	}
	return tenant, nil
}

func (s *Service) CreatTenant(t model.Tenant) (*model.Tenant, error) {
	if _, err := s.GetClient(t.ClientID); err != nil {
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
	if err := s.db.Where("id = ? AND client_id = ?", t.ID, t.ClientID).Save(&t).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteTenant(clientId, tenantId string) error {
	if _, err := s.GetTenant(clientId, tenantId); err != nil {
		return err
	}
	if err := s.db.Where("id = ? AND client_id = ?", tenantId, clientId).Delete(&model.Tenant{}).Error; err != nil {
		return err
	}
	return nil
}
