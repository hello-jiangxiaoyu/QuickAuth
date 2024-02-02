package oauth

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/pkg/global"
	"QuickAuth/pkg/safe"
	"errors"
	"go.uber.org/zap"
	"time"
)

var (
	ErrorCodeExpired = errors.New("code expired")
)

func GetAccessCode(appId string, codeName string) (*model.Code, error) {
	var code model.Code
	if err := global.Db().Where("app_id = ? AND code = ?", appId, codeName).
		First(&code).Error; err != nil {
		return nil, err
	}

	if err := global.Db().Where("id = ?", code.ID).Delete(model.Code{}).Error; err != nil {
		global.Log.Error("clear code err: ", zap.Error(err))
	}

	if code.CreatedAt.After(time.Now()) {
		return nil, ErrorCodeExpired
	}

	return &code, nil
}

func CreateAccessCode(appId string, userId string) (string, error) {
	code := safe.RandHex(31)
	state := safe.RandHex(31)
	accessCode := model.Code{
		AppID:  appId,
		UserID: userId,
		Code:   code,
		State:  state,
	}
	if err := global.Db().Create(accessCode).Error; err != nil {
		return "", err
	}

	return code, nil
}
