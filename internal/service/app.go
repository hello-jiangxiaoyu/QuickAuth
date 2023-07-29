package service

import (
	"QuickAuth/pkg/model"
	"QuickAuth/pkg/safe"
	"QuickAuth/pkg/utils"
	"errors"
	"github.com/lib/pq"
)

var ErrorDeleteDefaultApp = errors.New("do not delete the default app")

func (s *Service) ListApps() ([]model.App, error) {
	var apps []model.App
	if err := s.db.Select("id", "name", "icon", "tag").Find(&apps).Error; err != nil {
		return nil, err
	}

	return apps, nil
}

func (s *Service) GetApp(id string) (*model.App, error) {
	var app model.App
	if err := s.db.Where("id = ?", id).First(&app).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

func (s *Service) CreateApp(app model.App, host string, poolId int64) (*model.App, error) {
	app.ID = utils.GetNoLineUUID()
	if err := s.db.Create(&app).Error; err != nil {
		return nil, err
	}
	t := model.Tenant{
		AppID:        app.ID,
		UserPoolID:   poolId,
		Type:         1,
		Name:         "default",
		Company:      "default",
		Host:         host,
		Describe:     "default tenant created by app",
		RedirectUris: pq.StringArray{"http://localhost"},
		GrantTypes:   pq.StringArray{},
	}
	if err := s.db.Select("app_id", "user_pool_id", "type", "name", "host", "company", "grant_types", "redirect_uris", "describe").
		Create(&t).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

func (s *Service) ModifyApp(appId string, app model.App) error {
	if err := s.db.Where("id = ?", appId).Updates(&app).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteApp(appId string) error {
	if app, err := s.GetApp(appId); err != nil {
		return err
	} else if app.Name == "default" {
		return ErrorDeleteDefaultApp
	}
	if err := s.db.Where("id = ?", appId).Delete(model.App{}).Error; err != nil {
		return err
	}
	return nil
}

// =================== app secret ===================

func (s *Service) ListAppSecrets(appId string) ([]model.AppSecret, error) {
	var secrets []model.AppSecret
	if err := s.db.Where("app_id = ?", appId).Find(&secrets).Error; err != nil {
		return nil, err
	}

	return secrets, nil
}

func (s *Service) CreateAppSecret(secret model.AppSecret) (*model.AppSecret, error) {
	if _, err := s.GetApp(secret.AppID); err != nil {
		return nil, err
	}
	secret.Secret = safe.Rand62(63)
	if err := s.db.Create(&secret).Error; err != nil {
		return nil, err
	}
	return &secret, nil
}

func (s *Service) ModifyAppSecret(secretId int64, secret model.AppSecret) (*model.AppSecret, error) {
	if err := s.db.Select("scope", "access_expire", "refresh_expire", "describe").
		Where("id = ? AND app_id = ?", secretId, secret.AppID).Updates(&secret).Error; err != nil {
		return nil, err
	}
	return &secret, nil
}

func (s *Service) DeleteAppSecret(appId string, secretId int64) error {
	if err := s.db.Where("id = ? AND app_id = ?", secretId, appId).
		Delete(&model.AppSecret{}).Error; err != nil {
		return err
	}
	return nil
}
