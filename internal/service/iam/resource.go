package iam

import (
	"QuickAuth/pkg/conf"
	"QuickAuth/pkg/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ServiceIam struct {
	log  *zap.Logger
	db   *gorm.DB
	conf *conf.SystemConfig
}

func NewIamService(repo *global.Repository) *ServiceIam {
	return &ServiceIam{
		log:  repo.Logger,
		db:   repo.DB,
		conf: repo.Config,
	}
}

func (s *ServiceIam) ListResources() (interface{}, error) {
	var data interface{}
	// todo: add your service code hear!

	return data, nil
}

func (s *ServiceIam) GetResource() (interface{}, error) {
	var data interface{}
	// todo: add your service code hear!

	return data, nil
}

func (s *ServiceIam) CreateResource() (interface{}, error) {
	var data interface{}
	// todo: add your service code hear!

	return data, nil
}

func (s *ServiceIam) UpdateResource() error {
	// todo: add your service code hear!

	return nil
}

func (s *ServiceIam) DeleteResource() error {
	// todo: add your service code hear!

	return nil
}
