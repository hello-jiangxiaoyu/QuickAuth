package service

import "QuickAuth/pkg/model"

func (s *Service) GetTenant(id, clientId string) (*model.Tenant, error) {
	var tenant model.Tenant
	if err := s.db.Where("id = ? AND client_id = ?", id, clientId).
		First(&tenant).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (s *Service) ListTenant(clientId string) ([]model.Tenant, error) {
	var tenant []model.Tenant
	if err := s.db.Select("id", "client_id", "type", "name").Where("client_id = ?", clientId).Find(&tenant).Error; err != nil {
		return nil, err
	}
	return tenant, nil
}

func (s *Service) CreatTenant(t model.Tenant) (*model.Tenant, error) {
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

func (s *Service) DeleteTenant(t model.Tenant) (*model.Tenant, error) {
	if _, err := s.GetTenant(t.ID, t.ClientID); err != nil {
		return nil, err
	}
	if err := s.db.Where("id = ? AND client_id = ?", t.ID, t.ClientID).Delete(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}
