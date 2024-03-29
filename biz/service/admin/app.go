package admin

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/pkg/global"
	"QuickAuth/pkg/safe"
	"QuickAuth/pkg/utils"
	"errors"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

var ErrorDeleteDefaultApp = errors.New("do not delete the default app")

func ListApps() ([]model.App, error) {
	var apps []model.App
	if err := global.Db().Select("id", "name", "icon", "tag").Find(&apps).Error; err != nil {
		return nil, err
	}

	return apps, nil
}

func GetApp(id string) (*model.App, error) {
	var app model.App
	if err := global.Db().Where("id = ?", id).First(&app).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

func GetAppDetail(id string) (*model.App, error) {
	var app model.App
	if err := global.Db().Where("id = ?", id).Preload("Tenant").First(&app).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

func CreateApp(app *model.App, host string, poolId int64) (*model.App, error) {
	app.ID = utils.GetNoLineUUID()
	err := global.Db().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(app).Error; err != nil {
			return utils.WithMessage(err, "create app err")
		}

		userPool := model.UserPool{
			Name:     app.Name,
			Describe: app.Name + " user pool",
		}
		if poolId == 0 {
			if err := tx.Create(&userPool).Error; err != nil {
				return utils.WithMessage(err, "create default user pool err")
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
			return utils.WithMessage(err, "create default tenant err")
		}
		if err := tx.Create(&model.AppSecret{
			AppID:    app.ID,
			Secret:   safe.Rand62(63),
			Scope:    []string{"read_user"},
			Describe: "default secret",
		}).Error; err != nil {
			return utils.WithMessage(err, "create default secret err")
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return app, nil
}

func ModifyApp(appId string, app *model.App) error {
	if err := global.Db().Where("id = ?", appId).Updates(app).Error; err != nil {
		return err
	}
	return nil
}

func DeleteApp(appId string) error {
	if app, err := GetApp(appId); err != nil {
		return err
	} else if app.Name == "default" {
		return ErrorDeleteDefaultApp
	}

	var tenants []string
	if err := global.Db().Model(&model.Tenant{}).Select("id").Find(&tenants).Error; err != nil {
		return utils.WithMessage(err, "get tenant id list err")
	}

	err := global.Db().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("app_id = ?", appId).Delete(model.Tenant{}).Error; err != nil {
			return utils.WithMessage(err, "delete tenant err")
		}
		if err := tx.Where("app_id = ?", appId).Delete(model.Provider{}).Error; err != nil {
			return utils.WithMessage(err, "delete provider err")
		}
		if err := tx.Where("id = ?", appId).Delete(model.App{}).Error; err != nil {
			return utils.WithMessage(err, "delete app err")
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// =================== app secret ===================

func ListAppSecrets(appId string) ([]model.AppSecret, error) {
	var secrets []model.AppSecret
	if err := global.Db().Where("app_id = ?", appId).Find(&secrets).Error; err != nil {
		return nil, err
	}

	return secrets, nil
}

func CreateAppSecret(appId string, secret model.AppSecret) (*model.AppSecret, error) {
	if _, err := GetApp(appId); err != nil {
		return nil, err
	}
	secret.AppID = appId
	secret.Secret = safe.Rand62(63)
	if err := global.Db().Create(&secret).Error; err != nil {
		return nil, err
	}
	return &secret, nil
}

func ModifyAppSecret(secretId int64, secret model.AppSecret) (*model.AppSecret, error) {
	if err := global.Db().Select("scope", "access_expire", "refresh_expire", "describe").
		Where("id = ? AND app_id = ?", secretId, secret.AppID).Updates(&secret).Error; err != nil {
		return nil, err
	}
	return &secret, nil
}

func DeleteAppSecret(appId string, secretId int64) error {
	if err := global.Db().Where("id = ? AND app_id = ?", secretId, appId).
		Delete(&model.AppSecret{}).Error; err != nil {
		return err
	}
	return nil
}
