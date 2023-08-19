package internal

import (
	"QuickAuth/internal/endpoint/model"
	"QuickAuth/internal/service"
	"QuickAuth/pkg/conf"
	"QuickAuth/pkg/global"
	"QuickAuth/pkg/log"
	"QuickAuth/pkg/orm"
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

	db, err := orm.NewGormDB(global.Config.DB.DbType, dsn)
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
	if global.DB == nil {
		return errors.New("gorm db is not initialized")
	}
	if err := global.DB.Where("name = ?", "default").First(&global.App).Error; err == nil {
		return global.DB.Where("app_id = ?", global.App.ID).First(&global.Tenant).Error
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	svc := service.NewService(global.NewRepository(global.DB, global.Log, global.Config))
	app := &model.App{
		Name:     "default",
		Tag:      "Single Tenant",
		Icon:     "IconSafe",
		Describe: "quick auth app",
	}
	app, err := svc.CreateApp(app, "127.0.0.1", 0)
	if err != nil {
		return err
	}

	global.App = *app
	if err = global.DB.Where("app_id = ?", global.App.ID).First(&global.Tenant).Error; err != nil {
		return err
	}

	return nil
}
