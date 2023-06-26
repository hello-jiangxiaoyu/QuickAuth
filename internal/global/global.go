package global

import (
	"QuickAuth/pkg/conf"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	Config    *conf.SystemConfig
	Log       *zap.Logger
	AccessLog *zap.Logger
)

type Repository struct {
	// rdb    *redis.Client
	db     *gorm.DB
	logger *zap.Logger
	config *conf.SystemConfig
}

func NewRepository(db *gorm.DB, logger *zap.Logger, config *conf.SystemConfig) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
		config: config,
	}
}
