package service

import (
	"QuickAuth/internal/global"
	"QuickAuth/pkg/conf"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	log  *zap.Logger
	db   *gorm.DB
	conf *conf.SystemConfig
}

func NewService(repo *global.Repository) *Service {
	return &Service{
		log:  repo.Logger,
		db:   repo.DB,
		conf: repo.Config,
	}
}
