package service

import (
	"QuickAuth/pkg/model"
	"github.com/pkg/errors"
)

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
	if err := s.db.Model(model.Provider{}).Select("type", "app_id", "create_time").
		Where("tenant_id = ?", tenantId).Find(&provider).Error; err != nil {
		return nil, err
	}

	return provider, nil
}

func (s *Service) GetProvider(tenantId int64, providerId int64) (*model.Provider, error) {
	var provider model.Provider
	if err := s.db.Model(model.Provider{}).
		Where("id = ? AND tenant_id = ?", providerId, tenantId).First(&provider).Error; err != nil {
		return nil, err
	}

	return &provider, nil
}

func (s *Service) CreateProvider(prd model.Provider) (*model.Provider, error) {
	if err := s.db.Where("id = ?", prd.TenantID).First(&model.Tenant{}).Error; err != nil {
		return nil, errors.Wrap(err, "no such tenant")
	}
	if err := s.db.Create(&prd).Error; err != nil {
		return nil, err
	}
	return &prd, nil
}

func (s *Service) ModifyProvider(prd model.Provider) error {
	if err := s.db.Where("id = ? ANT tenant_id = ?", prd.ID, prd.TenantID).
		Save(&prd).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteProvider(tenantId int64, providerId int64) error {
	if err := s.db.Where("id = ? ANT tenant_id = ?", providerId, tenantId).
		Delete(model.Provider{}).Error; err != nil {
		return err
	}
	return nil
}
