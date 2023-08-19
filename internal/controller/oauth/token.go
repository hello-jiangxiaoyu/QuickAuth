package oauth

import (
	"QuickAuth/internal/endpoint/dto"
	"QuickAuth/internal/endpoint/request"
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

func (o Controller) getTokenHandler(grantType string) (Handler, error) {
	var tokenHandler = map[string]Handler{
		grantTypeCode:         o.handlerAuthorizationCode,
		grantTypeClient:       o.handlerClientCredential,
		grantTypePassword:     o.handlerPassword,
		grantTypeRefreshToken: o.handlerRefreshToken,
	}
	h, ok := tokenHandler[grantType]
	if !ok {
		return nil, errors.New("invalid grant_type")
	}
	return h, nil
}

func (o Controller) handlerAuthorizationCode(req *request.Token) (*dto.TokenResponse, error) {
	code, err := o.svc.GetAccessCode(req.ClientID, req.Code)
	if err != nil {
		return nil, err
	}

	token, err := o.svc.CreateAccessToken(req.App, req.Tenant, req.UserID, req.Nonce, code.Scope)
	if err != nil {
		return nil, err
	}
	res := &dto.TokenResponse{AccessToken: token}
	return res, nil
}

func (o Controller) handlerClientCredential(req *request.Token) (*dto.TokenResponse, error) {
	token, err := o.svc.CreateAccessToken(req.App, req.Tenant, req.UserID, req.Nonce, []string{req.State})
	if err != nil {
		return nil, err
	}

	res := &dto.TokenResponse{AccessToken: token}
	return res, nil
}

func (o Controller) handlerPassword(req *request.Token) (*dto.TokenResponse, error) {
	token, err := o.svc.CreateAccessToken(req.App, req.Tenant, req.UserID, req.Nonce, []string{req.State})
	if err != nil {
		return nil, err
	}

	res := &dto.TokenResponse{AccessToken: token}
	return res, nil
}

func (o Controller) handlerRefreshToken(req *request.Token) (*dto.TokenResponse, error) {
	token, err := o.svc.CreateAccessToken(req.App, req.Tenant, req.UserID, req.Nonce, []string{req.State})
	if err != nil {
		return nil, err
	}

	res := &dto.TokenResponse{AccessToken: token}
	return res, nil
}
