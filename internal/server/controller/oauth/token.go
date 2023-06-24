package oauth

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/model"
	"QuickAuth/internal/server/service"
	"errors"
)

const (
	grantTypeCode   = "authorization_code"
	grantTypeClient = "client_credential"
)

type TokenResponse struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireIn     int    `json:"expire_in"`
	ExpireAt     int    `json:"expire_at"`
}
type Handler func(*request.Token) (*TokenResponse, error)

var tokenHandler = map[string]Handler{
	grantTypeCode:   authorizationCodeHandler,
	grantTypeClient: clientCredentialHandler,
}

func getTokenHandler(grantType string) (Handler, error) {
	h, ok := tokenHandler[grantType]
	if !ok {
		return nil, errors.New("invalid grant_type")
	}
	return h, nil
}

func authorizationCodeHandler(req *request.Token) (*TokenResponse, error) {
	token, err := service.CreateAccessToken(model.Client{ID: req.ClientID},
		req.Tenant.ID, req.Tenant.Host, req.Nonce, req.State, req.UserID)
	if err != nil {
		return nil, err
	}
	res := &TokenResponse{AccessToken: token}
	return res, nil
}

func clientCredentialHandler(req *request.Token) (*TokenResponse, error) {
	token, err := service.CreateAccessToken(model.Client{ID: req.ClientID},
		req.Tenant.ID, req.Tenant.Host, req.Nonce, req.State, req.UserID)
	if err != nil {
		return nil, err
	}

	res := &TokenResponse{AccessToken: token}
	return res, nil
}
