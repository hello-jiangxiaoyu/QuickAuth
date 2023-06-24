package service

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/global"
	"QuickAuth/internal/model"
	"QuickAuth/internal/utils"
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

func IsClientValid(clientId string) bool {
	var client model.Client
	if err := global.DB.Select("id").Where("id = ?", clientId).First(&client).Error; err != nil {
		global.Log.Error("get secret err: ", zap.Error(err))
		return false
	}

	return true
}

func IsClientSecretValid(clientId, clientSecret string) bool {
	var secret model.ClientSecret
	if err := global.DB.Select("secret").
		Where("client_id = ? AND secret = ?", clientId, clientSecret).
		First(&secret).Error; err != nil {
		global.Log.Error("get secret err: ", zap.Error(err))
		return false
	}

	return secret.Secret == clientSecret
}

func IsCodeValid(clientId string, codeName string) bool {
	var code model.Code
	if err := global.DB.Select("id").
		Where("client_id = ? AND code = ?", clientId, codeName).
		Find(&code).Error; err != nil {
		global.Log.Error("get code err: ", zap.Error(err))
		return false
	}

	if err := global.DB.Where("id = ?", code.ID).Delete(model.Code{}).Error; err != nil {
		global.Log.Error("clear code err: ", zap.Error(err))
	}
	return true
}

func IsRedirectUriValid(clientId, uri string) bool {
	var redirectUri model.RedirectURI
	if err := global.DB.Select("uri").
		Where("client_id = ? AND uri = ?", clientId, uri).
		First(&uri).Error; err != nil {
		global.Log.Error("failed to find redirect uri: ", zap.Error(err))
		return false
	}
	return redirectUri.URI == uri
}

func CreateAccessCode(clientId string, userId string) (string, string, error) {
	code := utils.RandHex(31)
	state := utils.RandHex(31)
	accessCode := model.Code{
		ClientID: clientId,
		UserID:   userId,
		Code:     code,
		State:    state,
	}
	if err := global.DB.Create(accessCode).Error; err != nil {
		return "", "", err
	}

	return code, state, nil
}
