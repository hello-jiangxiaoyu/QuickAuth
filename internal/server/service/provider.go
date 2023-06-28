package service

import "QuickAuth/pkg/model"

func (s *Service) GetProviderByType(tenantId, name string) (*model.Provider, error) {
	var provider model.Provider
	if err := s.db.Model(model.Provider{}).
		Where("tenant_id = ? AND type = ?", tenantId, name).First(&provider).Error; err != nil {
		return nil, err
	}

	return &provider, nil
}

func (s *Service) GetLoginProviderInfo(tenantId string) ([]model.Provider, error) {
	var provider []model.Provider
	if err := s.db.Model(model.Provider{}).Select("type", "client_id").
		Where("tenant_id = ?", tenantId).Find(&provider).Error; err != nil {
		return nil, err
	}

	return provider, nil
}
