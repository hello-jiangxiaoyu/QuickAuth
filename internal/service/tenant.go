package service

import (
	"QuickAuth/pkg/model"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"net/url"
)

func (s *Service) GetTenant(appId string, tenantId int64) (*model.Tenant, error) {
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
	if _, err := s.GetApp(t.AppID); err != nil {
		return nil, errors.Wrap(err, "no such app")
	}
	if _, err := s.GetUserPool(t.UserPoolID); err != nil {
		return nil, errors.Wrap(err, "no such user pool")
	}

	t.RedirectUris = pq.StringArray{"https://" + t.Host, "http://" + t.Host}
	t.GrantTypes = pq.StringArray{"authorization_code", "client_credential", "refresh_token"}
	if err := s.db.Select("app_id", "user_pool_id", "type", "name", "host", "company", "grant_types", "redirect_uris", "describe").
		Create(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *Service) ModifyTenant(tenantId int64, t model.Tenant) error {
	if err := s.db.Select("app_id", "user_pool_id", "type", "name", "host", "company", "grant_type", "describe").
		Where("id = ? AND app_id = ?", tenantId, t.AppID).Updates(&t).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteTenant(appId string, tenantId int64) error {
	if err := s.db.Where("id = ? AND app_id = ?", tenantId, appId).Delete(&model.Tenant{}).Error; err != nil {
		return err
	}
	return nil
}

// =================== redirect uri ===================

func (s *Service) IsRedirectUriValid(appId string, tenantId int64, uri string) (bool, error) {
	var app model.Tenant
	if err := s.db.Select("uri").Where("id = ? AND app_id = ?", tenantId, appId).First(&app).Error; err != nil {
		return false, err
	}

	for _, v := range app.RedirectUris {
		if v == uri {
			return true, nil
		}
	}
	return false, nil
}

func (s *Service) ListRedirectUri(appId string, tenantId int64) ([]string, error) {
	var apps []string
	if err := s.db.Model(model.Tenant{}).Select("redirect_uris").
		Where("id = ? AND app_id = ?", tenantId, appId).Find(&apps).Error; err != nil {
		return nil, err
	}

	return apps, nil
}

func (s *Service) CreateRedirectUri(appId string, tenantId int64, uri string) error {
	sql := `update tenants set redirect_uris = array_prepend(?, redirect_uris) where id = ? and app_id = ?;`
	if err := s.db.Exec(sql, uri, tenantId, appId).Error; err != nil {
		return err
	}

	return nil
}

func (s *Service) ModifyRedirectUri(appId string, tenantId int64, uriId uint, uri string) error {
	sql := `update tenants set redirect_uris[?] = ? where id = ? and app_id = ?;`
	if err := s.db.Exec(sql, uriId, uri, tenantId, appId).Error; err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteRedirectUri(appId string, tenantId int64, uri string) error {
	uri, err := url.QueryUnescape(uri)
	if err != nil {
		return errors.Wrap(err, "invalid uri")
	}

	sql := `update tenants set redirect_uris = array_remove(redirect_uris, ?) where id = ? and app_id = ?;`
	if err = s.db.Exec(sql, uri, tenantId, appId).Error; err != nil {
		return err
	}

	return nil
}
