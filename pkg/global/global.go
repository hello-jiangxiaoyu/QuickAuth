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

type Repository struct {
	// rdb    *redis.Client
	DB     *gorm.DB
	Logger *zap.Logger
	Config *conf.SystemConfig
}

func Db() *gorm.DB {
	return DB
}

func NewRepository(db *gorm.DB, logger *zap.Logger, config *conf.SystemConfig) *Repository {
	return &Repository{
		DB:     db,
		Logger: logger,
		Config: config,
	}
}
