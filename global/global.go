package global

import (
	"QuickAuth/conf"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	Config    *conf.SystemConfig
	Log       *zap.Logger
	AccessLog *zap.Logger
)
