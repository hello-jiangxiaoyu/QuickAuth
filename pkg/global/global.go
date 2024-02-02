package global

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/pkg/conf"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	Config    *conf.SystemConfig
	Log       *zap.Logger
	AccessLog *zap.Logger
	App       model.App
	Tenant    model.Tenant
)

func Db() *gorm.DB {
	return DB
}
