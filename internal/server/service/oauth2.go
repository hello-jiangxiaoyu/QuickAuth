package service

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/global"
	"QuickAuth/internal/server/model"
	"go.uber.org/zap"
)

func GetUser(req *request.Login) (*model.User, error) {
	var user model.User
	if err := global.DB.Where("user_pool_id = ? AND username = ?", req.Tenant.UserPoolID, req.UserName).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func IsClientSecretValid(clientId, clientSecret string) bool {
	secrets := make([]model.ClientSecret, 0)
	if err := global.DB.Where("client_id = ?", clientId).Find(&secrets).Error; err != nil {
		global.Log.Error("get secret err: ", zap.Error(err))
		return false
	}
	for _, secret := range secrets {
		if *secret.Secret == clientSecret {
			return true
		}
	}
	return false
}

func IsCodeValid(clientId string, codeName string) bool {
	var code model.Code
	if err := global.DB.Where("client_id = ? AND code = ?", clientId, codeName).
		Find(&code).Error; err != nil {
		global.Log.Error("get code err: ", zap.Error(err))
		return false
	}

	if err := global.DB.Where("id = ?", code.ID).Delete(model.Code{}).Error; err != nil {
		global.Log.Error("clear code err: ", zap.Error(err))
	}
	return true
}
