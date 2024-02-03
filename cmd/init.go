package cmd

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/cmd/log"
	"QuickAuth/pkg/conf"
	"QuickAuth/pkg/global"
	"QuickAuth/pkg/safe"
	"QuickAuth/pkg/utils"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

func InitConfig(path string) error {
	if global.Config != nil { // 初始化过一次
		fmt.Println("global.Config is already initialized")
		return nil
	}
	c, err := conf.NewSystemConfig(path)
	if err != nil {
		return err
	}
	global.Config = c
	return nil
}

func InitGorm() error {
	if global.Config == nil {
		return errors.New("global.Config is nil, failed to initialize gorm")
	}
	dsn := global.Config.DB.GetDsn()
	if dsn == "" {
		return errors.New("dsn is empty, failed to initialize gorm")
	}

	db, err := global.NewGormDB(dsn)
	if err != nil {
		return err
	}
	global.DB = db
	return nil
}

func InitLogger() error {
	if global.Config == nil {
		return errors.New("global.Config is nil, failed to initialize logger")
	}
	global.Log = log.NewZapErrorLogger(global.Config.Log.Dir)
	global.AccessLog = log.NewZapAccessLogger(global.Config.Log.Dir)
	return nil
}

func InitDefaultTenant() error {
	if global.Db() == nil {
		return errors.New("gorm db is not initialized")
	}
	if err := global.Db().Where("name = ?", "default").First(&global.App).Error; err == nil {
		return global.Db().Where("app_id = ?", global.App.ID).First(&global.Tenant).Error
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	app := &model.App{
		ID:       utils.GetNoLineUUID(),
		Name:     "default",
		Tag:      "Single Tenant",
		Icon:     "IconSafe",
		Describe: "quick auth app",
	}
	t := model.Tenant{
		AppID:        app.ID,
		Type:         1,
		Name:         "default",
		Company:      "default",
		Host:         "127.0.0.1:8000",
		Describe:     "default tenant created by app",
		RedirectUris: pq.StringArray{"http://127.0.0.1:8000"},
		GrantTypes:   pq.StringArray{},
		Config:       "{}",
	}

	err := global.Db().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(app).Error; err != nil {
			return utils.WithMessage(err, "create app err")
		}
		userPool := model.UserPool{
			Name:     app.Name,
			Describe: app.Name + " user pool",
		}
		if err := tx.Create(&userPool).Error; err != nil {
			return utils.WithMessage(err, "create default user pool err")
		}
		t.UserPoolID = userPool.ID
		if err := tx.Create(&model.User{
			UserPoolID:  userPool.ID,
			Username:    "admin",
			Password:    safe.HashPassword("admin"),
			DisplayName: "admin",
		}).Error; err != nil {
			return utils.WithMessage(err, "create user err")
		}
		if err := tx.Create(&t).Error; err != nil {
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
		return err
	}

	global.App = *app
	if err = global.Db().Where("app_id = ?", global.App.ID).First(&global.Tenant).Error; err != nil {
		return err
	}

	return nil
}
