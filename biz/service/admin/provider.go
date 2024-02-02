package admin

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/pkg/global"
)

func GetProviderByType(tenantId int64, name string) (*model.Provider, error) {
	var provider model.Provider
	if err := global.Db().Model(model.Provider{}).
		Where("tenant_id = ? AND type = ?", tenantId, name).First(&provider).Error; err != nil {
		return nil, err
	}

	return &provider, nil
}

func GetLoginProviderInfo(tenantId int64) ([]model.Provider, error) {
	var provider []model.Provider
	if err := global.Db().Model(model.Provider{}).Select("id", "type", "tenant_id", "client_id", "created_at").
		Where("tenant_id = ?", tenantId).Find(&provider).Error; err != nil {
		return nil, err
	}

	return provider, nil
}

func GetProviderById(tenantId int64, providerId int64) (*model.Provider, error) {
	var provider model.Provider
	if err := global.Db().Model(model.Provider{}).
		Where("id = ? AND tenant_id = ?", providerId, tenantId).First(&provider).Error; err != nil {
		return nil, err
	}

	return &provider, nil
}

func CreateProvider(prd *model.Provider) (*model.Provider, error) {
	prd.ID = 0
	if err := global.Db().Create(prd).Error; err != nil {
		return nil, err
	}
	return prd, nil
}

func ModifyProvider(providerId int64, prd *model.Provider) error {
	if err := global.Db().Where("id = ? AND tenant_id = ?", providerId, prd.TenantID).
		Updates(prd).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProvider(tenantId int64, providerId int64) error {
	if err := global.Db().Where("id = ? AND tenant_id = ?", providerId, tenantId).
		Delete(model.Provider{}).Error; err != nil {
		return err
	}
	return nil
}
