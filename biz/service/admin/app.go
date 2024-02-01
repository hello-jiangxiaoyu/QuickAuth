package admin

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/pkg/conf"
	"QuickAuth/pkg/global"
	"QuickAuth/pkg/safe"
	"QuickAuth/pkg/utils"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var ErrorDeleteDefaultApp = errors.New("do not delete the default app")

type ServiceAdmin struct {
	log  *zap.Logger
	db   *gorm.DB
	conf *conf.SystemConfig
}

func NewAdminService(repo *global.Repository) *ServiceAdmin {
	return &ServiceAdmin{
		log:  repo.Logger,
		db:   repo.DB,
		conf: repo.Config,
	}
}

func (s *ServiceAdmin) ListApps() ([]model.App, error) {
	var apps []model.App
	if err := s.db.Select("id", "name", "icon", "tag").Find(&apps).Error; err != nil {
		return nil, err
	}

	return apps, nil
}

func (s *ServiceAdmin) GetApp(id string) (*model.App, error) {
	var app model.App
	if err := s.db.Where("id = ?", id).First(&app).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

func (s *ServiceAdmin) GetAppDetail(id string) (*model.App, error) {
	var app model.App
	if err := s.db.Where("id = ?", id).Preload("Tenant").First(&app).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

func (s *ServiceAdmin) CreateApp(app *model.App, host string, poolId int64) (*model.App, error) {
	app.ID = utils.GetNoLineUUID()
	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(app).Error; err != nil {
			return errors.WithMessage(err, "create app err")
		}

		userPool := model.UserPool{
			Name:     app.Name,
			Describe: app.Name + " user pool",
		}
		if poolId == 0 {
			if err := tx.Create(&userPool).Error; err != nil {
				return errors.WithMessage(err, "create default user pool err")
			}
			poolId = userPool.ID
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
			Config:       "{}",
		}
		if err := tx.Select("app_id", "user_pool_id", "type", "name", "host", "company", "grant_types",
			"redirect_uris", "describe", "config").Create(&t).Error; err != nil {
			return errors.WithMessage(err, "create default tenant err")
		}
		if err := tx.Create(&model.AppSecret{
			AppID:    app.ID,
			Secret:   safe.Rand62(63),
			Scope:    []string{"read_user"},
			Describe: "default secret",
		}).Error; err != nil {
			return errors.WithMessage(err, "create default secret err")
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return app, nil
}

func (s *ServiceAdmin) ModifyApp(appId string, app *model.App) error {
	if err := s.db.Where("id = ?", appId).Updates(app).Error; err != nil {
		return err
	}
	return nil
}

func (s *ServiceAdmin) DeleteApp(appId string) error {
	if app, err := s.GetApp(appId); err != nil {
		return err
	} else if app.Name == "default" {
		return ErrorDeleteDefaultApp
	}

	var tenants []string
	if err := s.db.Model(&model.Tenant{}).Select("id").Find(&tenants).Error; err != nil {
		return errors.WithMessage(err, "get tenant id list err")
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("app_id = ?", appId).Delete(model.Tenant{}).Error; err != nil {
			return errors.WithMessage(err, "delete tenant err")
		}
		if err := tx.Where("app_id = ?", appId).Delete(model.Provider{}).Error; err != nil {
			return errors.WithMessage(err, "delete provider err")
		}
		if err := tx.Where("id = ?", appId).Delete(model.App{}).Error; err != nil {
			return errors.WithMessage(err, "delete app err")
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// =================== app secret ===================

func (s *ServiceAdmin) ListAppSecrets(appId string) ([]model.AppSecret, error) {
	var secrets []model.AppSecret
	if err := s.db.Where("app_id = ?", appId).Find(&secrets).Error; err != nil {
		return nil, err
	}

	return secrets, nil
}

func (s *ServiceAdmin) CreateAppSecret(appId string, secret model.AppSecret) (*model.AppSecret, error) {
	if _, err := s.GetApp(appId); err != nil {
		return nil, err
	}
	secret.AppID = appId
	secret.Secret = safe.Rand62(63)
	if err := s.db.Create(&secret).Error; err != nil {
		return nil, err
	}
	return &secret, nil
}

func (s *ServiceAdmin) ModifyAppSecret(secretId int64, secret model.AppSecret) (*model.AppSecret, error) {
	if err := s.db.Select("scope", "access_expire", "refresh_expire", "describe").
		Where("id = ? AND app_id = ?", secretId, secret.AppID).Updates(&secret).Error; err != nil {
		return nil, err
	}
	return &secret, nil
}

func (s *ServiceAdmin) DeleteAppSecret(appId string, secretId int64) error {
	if err := s.db.Where("id = ? AND app_id = ?", secretId, appId).
		Delete(&model.AppSecret{}).Error; err != nil {
		return err
	}
	return nil
}