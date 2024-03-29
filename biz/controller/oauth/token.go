package oauth

import (
	"QuickAuth/biz/endpoint/dto"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/service/oauth"
	"errors"
)

const (
	grantTypeCode         = "authorization_code"
	grantTypeClient       = "client_credential"
	grantTypePassword     = "password"
	grantTypeRefreshToken = "refresh_token"
	grantTypeDeviceFlow   = "device_code"
)

type Handler func(*request.Token) (*dto.TokenResponse, error)

func getTokenHandler(grantType string) (Handler, error) {
	var tokenHandler = map[string]Handler{
		grantTypeCode:         handlerAuthorizationCode,
		grantTypeClient:       handlerClientCredential,
		grantTypePassword:     handlerPassword,
		grantTypeRefreshToken: handlerRefreshToken,
	}
	h, ok := tokenHandler[grantType]
	if !ok {
		return nil, errors.New("invalid grant_type")
	}
	return h, nil
}

func handlerAuthorizationCode(req *request.Token) (*dto.TokenResponse, error) {
	code, err := oauth.GetAccessCode(req.ClientID, req.Code)
	if err != nil {
		return nil, err
	}

	token, err := oauth.CreateAccessToken(req.App, req.Tenant, req.UserID, req.Nonce, code.Scope)
	if err != nil {
		return nil, err
	}
	res := &dto.TokenResponse{AccessToken: token}
	return res, nil
}

func handlerClientCredential(req *request.Token) (*dto.TokenResponse, error) {
	token, err := oauth.CreateAccessToken(req.App, req.Tenant, req.UserID, req.Nonce, []string{req.State})
	if err != nil {
		return nil, err
	}

	res := &dto.TokenResponse{AccessToken: token}
	return res, nil
}

func handlerPassword(req *request.Token) (*dto.TokenResponse, error) {
	token, err := oauth.CreateAccessToken(req.App, req.Tenant, req.UserID, req.Nonce, []string{req.State})
	if err != nil {
		return nil, err
	}

	res := &dto.TokenResponse{AccessToken: token}
	return res, nil
}

func handlerRefreshToken(req *request.Token) (*dto.TokenResponse, error) {
	token, err := oauth.CreateAccessToken(req.App, req.Tenant, req.UserID, req.Nonce, []string{req.State})
	if err != nil {
		return nil, err
	}

	res := &dto.TokenResponse{AccessToken: token}
	return res, nil
}
