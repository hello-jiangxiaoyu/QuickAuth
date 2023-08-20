package iam

import (
	"QuickAuth/internal/endpoint/model"
	"QuickAuth/internal/endpoint/request"
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

func (s *ServiceIam) ListResources(tenantId int64) ([]model.Resource, error) {
	var data []model.Resource
	if err := s.db.Where("tenant_id = ?", tenantId).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (s *ServiceIam) GetResource(tenantId int64, in *request.Iam) (*model.Resource, error) {
	var data *model.Resource
	if err := s.db.Where("tenant_id = ?", tenantId, in.ResourceId).First(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (s *ServiceIam) CreateResource(res *model.Resource) (*model.Resource, error) {
	if err := s.db.Create(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ServiceIam) UpdateResource(tenantId int64, resId int64, res *model.Resource) error {
	if err := s.db.Where("tenant_id = ? AND id = ?", tenantId, resId).Updates(res).Error; err != nil {
		return err
	}

	return nil
}

func (s *ServiceIam) DeleteResource(tenantId int64, resId int64) error {
	if err := s.db.Where("tenant_id = ? AND id = ?", tenantId, resId).Delete(model.Resource{}).Error; err != nil {
		return err
	}

	return nil
}
