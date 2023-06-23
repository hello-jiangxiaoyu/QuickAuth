package request

import (
	"QuickAuth/internal/server/model"
)

type Login struct {
	UserName string       `json:"username" binding:"required"`
	Password string       `json:"password" binding:"required"`
	Tenant   model.Tenant `json:"-"`
}

type Auth struct {
	ClientId     string       `json:"clientId" binding:"required"`
	Scope        string       `json:"scope" binding:"required"`
	ResponseType string       `json:"responseType" binding:"required"`
	RedirectUri  string       `json:"redirectUri" binding:"required"`
	State        string       `json:"state" binding:"required"`
	Nonce        string       `json:"nonce"`
	Tenant       model.Tenant `json:"-"`
}

type Token struct {
	ClientId     string       `json:"clientId" binding:"required"`
	ClientSecret string       `json:"clientSecret" binding:"required"`
	GrantType    string       `json:"grantType" binding:"required"`
	Code         string       `json:"code" binding:"required"`
	RedirectUri  string       `json:"redirectUri" binding:"required"`
	State        string       `json:"state" binding:"required"`
	Nonce        string       `json:"nonce" binding:"required"`
	Tenant       model.Tenant `json:"-"`
}
