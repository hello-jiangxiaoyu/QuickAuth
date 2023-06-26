package internal

import (
	"QuickAuth/internal/global"
	"QuickAuth/pkg/conf"
	"QuickAuth/pkg/gorm"
	"QuickAuth/pkg/log"
	"errors"
)

func InitConfig() error {
	c, err := conf.NewSystemConfig()
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

	db, err := gorm.NewGormDB(global.Config.DB.DbType, dsn)
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
