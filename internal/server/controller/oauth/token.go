package oauth

import (
	"QuickAuth/internal/endpoint/request"
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
		return nil, errors.New("no such grant type")
	}
	return h, nil
}

func authorizationCodeHandler(req *request.Token) (*TokenResponse, error) {
	return nil, nil
}
func clientCredentialHandler(req *request.Token) (*TokenResponse, error) {
	return nil, nil
}
