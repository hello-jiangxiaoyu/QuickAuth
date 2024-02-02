package internal

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/biz/service/admin"
	"QuickAuth/pkg/conf"
	"QuickAuth/pkg/global"
	"QuickAuth/pkg/log"
	"errors"
	"fmt"
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
	errorLog, err := log.NewZapErrorLogger(global.Config.Log.Dir, global.Config.Log.Level)
	if err != nil {
		return err
	}
	accessLog, err := log.NewZapAccessLogger(global.Config.Log.Dir)
	if err != nil {
		return err
	}

	global.Log = errorLog
	global.AccessLog = accessLog
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
		Name:     "default",
		Tag:      "Single Tenant",
		Icon:     "IconSafe",
		Describe: "quick auth app",
	}
	app, err := admin.CreateApp(app, "127.0.0.1", 0)
	if err != nil {
		return err
	}

	global.App = *app
	if err = global.Db().Where("app_id = ?", global.App.ID).First(&global.Tenant).Error; err != nil {
		return err
	}

	return nil
}
