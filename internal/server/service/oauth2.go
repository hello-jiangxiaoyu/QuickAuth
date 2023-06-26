package service

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/global"
	"QuickAuth/pkg/model"
	"QuickAuth/pkg/utils/safe"
	"errors"
	"go.uber.org/zap"
	"time"
)

var (
	ErrorCodeExpired = errors.New("code expired")
)

func GetUser(req *request.Login) (*model.User, error) {
	var user model.User
	if err := global.DB.Where("user_pool_id = ? AND username = ?", req.Tenant.UserPoolID, req.UserName).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func IsClientValid(clientId string) (bool, error) {
	var client model.Client
	if err := global.DB.Select("id").Where("id = ?", clientId).Limit(1).Find(&client).Error; err != nil {
		return false, err
	}

	return client.ID == clientId, nil
}

func IsClientSecretValid(clientId, clientSecret string) (bool, error) {
	var secret model.ClientSecret
	if err := global.DB.Select("secret").
		Where("client_id = ? AND secret = ?", clientId, clientSecret).
		Limit(1).Find(&secret).Error; err != nil {
		return false, err
	}

	return secret.Secret == clientSecret, nil
}

func GetAccessCode(clientId string, codeName string) (*model.Code, error) {
	var code model.Code
	if err := global.DB.Where("client_id = ? AND code = ?", clientId, codeName).
		First(&code).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("id = ?", code.ID).Delete(model.Code{}).Error; err != nil {
		global.Log.Error("clear code err: ", zap.Error(err))
	}

	if code.CreateTime.After(time.Now()) {
		return nil, ErrorCodeExpired
	}

	return &code, nil
}

func IsRedirectUriValid(clientId, uri string) (bool, error) {
	var redirectUri model.RedirectURI
	if err := global.DB.Select("uri").Where("client_id = ? AND uri = ?", clientId, uri).
		Limit(1).Find(&uri).Error; err != nil {
		return false, err
	}

	return redirectUri.URI == uri, nil
}

func CreateAccessCode(clientId string, userId string) (string, string, error) {
	code := safe.RandHex(31)
	state := safe.RandHex(31)
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
