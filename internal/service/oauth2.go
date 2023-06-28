package service

import (
	"QuickAuth/pkg/model"
	"QuickAuth/pkg/tools/safe"
	"errors"
	"go.uber.org/zap"
	"time"
)

var (
	ErrorCodeExpired = errors.New("code expired")
)

func (s *Service) GetAccessCode(clientId string, codeName string) (*model.Code, error) {
	var code model.Code
	if err := s.db.Where("client_id = ? AND code = ?", clientId, codeName).
		First(&code).Error; err != nil {
		return nil, err
	}

	if err := s.db.Where("id = ?", code.ID).Delete(model.Code{}).Error; err != nil {
		s.log.Error("clear code err: ", zap.Error(err))
	}

	if code.CreateTime.After(time.Now()) {
		return nil, ErrorCodeExpired
	}

	return &code, nil
}

func (s *Service) CreateAccessCode(clientId string, userId string) (string, string, error) {
	code := safe.RandHex(31)
	state := safe.RandHex(31)
	accessCode := model.Code{
		ClientID: clientId,
		UserID:   userId,
		Code:     code,
		State:    state,
	}
	if err := s.db.Create(accessCode).Error; err != nil {
		return "", "", err
	}

	return code, state, nil
}
