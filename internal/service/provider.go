package service

import "QuickAuth/pkg/model"

func (s *Service) GetProviderByType(tenantId int64, name string) (*model.Provider, error) {
	var provider model.Provider
	if err := s.db.Model(model.Provider{}).
		Where("tenant_id = ? AND type = ?", tenantId, name).First(&provider).Error; err != nil {
		return nil, err
	}

	return &provider, nil
}

func (s *Service) GetLoginProviderInfo(tenantId int64) ([]model.Provider, error) {
	var provider []model.Provider
	if err := s.db.Model(model.Provider{}).Select("type", "client_id").
		Where("tenant_id = ?", tenantId).Find(&provider).Error; err != nil {
		return nil, err
	}

	return provider, nil
}

func (s *Service) CreateProvider(prd model.Provider) (*model.Provider, error) {
	// todo: check if tenant exists
	if err := s.db.Create(&prd).Error; err != nil {
		return nil, err
	}
	return &prd, nil
}

func (s *Service) ModifyProvider(prd model.Provider) error {
	if err := s.db.Where("tenant_id = ? ANT type = ?", prd.TenantID, prd.Type).
		Save(&prd).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteProvider(tenantId int64, providerName string) error {
	if err := s.db.Where("tenant_id = ? ANT type = ?", tenantId, providerName).
		Delete(model.Provider{}).Error; err != nil {
		return err
	}
	return nil
}
