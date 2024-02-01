package admin

import (
	"QuickAuth/biz/endpoint/model"
)

func (s *ServiceAdmin) GetProviderByType(tenantId int64, name string) (*model.Provider, error) {
	var provider model.Provider
	if err := s.db.Model(model.Provider{}).
		Where("tenant_id = ? AND type = ?", tenantId, name).First(&provider).Error; err != nil {
		return nil, err
	}

	return &provider, nil
}

func (s *ServiceAdmin) GetLoginProviderInfo(tenantId int64) ([]model.Provider, error) {
	var provider []model.Provider
	if err := s.db.Model(model.Provider{}).Select("id", "type", "tenant_id", "client_id", "created_at").
		Where("tenant_id = ?", tenantId).Find(&provider).Error; err != nil {
		return nil, err
	}

	return provider, nil
}

func (s *ServiceAdmin) GetProviderById(tenantId int64, providerId int64) (*model.Provider, error) {
	var provider model.Provider
	if err := s.db.Model(model.Provider{}).
		Where("id = ? AND tenant_id = ?", providerId, tenantId).First(&provider).Error; err != nil {
		return nil, err
	}

	return &provider, nil
}

func (s *ServiceAdmin) CreateProvider(prd *model.Provider) (*model.Provider, error) {
	prd.ID = 0
	if err := s.db.Create(prd).Error; err != nil {
		return nil, err
	}
	return prd, nil
}

func (s *ServiceAdmin) ModifyProvider(providerId int64, prd *model.Provider) error {
	if err := s.db.Where("id = ? AND tenant_id = ?", providerId, prd.TenantID).
		Updates(prd).Error; err != nil {
		return err
	}
	return nil
}

func (s *ServiceAdmin) DeleteProvider(tenantId int64, providerId int64) error {
	if err := s.db.Where("id = ? AND tenant_id = ?", providerId, tenantId).
		Delete(model.Provider{}).Error; err != nil {
		return err
	}
	return nil
}