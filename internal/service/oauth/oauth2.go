package oauth

import (
	"QuickAuth/internal/endpoint/model"
	"QuickAuth/pkg/conf"
	"QuickAuth/pkg/global"
	"QuickAuth/pkg/safe"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	ErrorCodeExpired = errors.New("code expired")
)

type ServiceOauth struct {
	log  *zap.Logger
	db   *gorm.DB
	conf *conf.SystemConfig
}

func NewOauthService(repo *global.Repository) *ServiceOauth {
	return &ServiceOauth{
		log:  repo.Logger,
		db:   repo.DB,
		conf: repo.Config,
	}
}

func (s *ServiceOauth) GetAccessCode(appId string, codeName string) (*model.Code, error) {
	var code model.Code
	if err := s.db.Where("app_id = ? AND code = ?", appId, codeName).
		First(&code).Error; err != nil {
		return nil, err
	}

	if err := s.db.Where("id = ?", code.ID).Delete(model.Code{}).Error; err != nil {
		s.log.Error("clear code err: ", zap.Error(err))
	}

	if code.CreatedAt.After(time.Now()) {
		return nil, ErrorCodeExpired
	}

	return &code, nil
}

func (s *ServiceOauth) CreateAccessCode(appId string, userId string) (string, error) {
	code := safe.RandHex(31)
	state := safe.RandHex(31)
	accessCode := model.Code{
		AppID:  appId,
		UserID: userId,
		Code:   code,
		State:  state,
	}
	if err := s.db.Create(accessCode).Error; err != nil {
		return "", err
	}

	return code, nil
}
