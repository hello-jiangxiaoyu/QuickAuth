package request

import (
	"QuickAuth/pkg/models"
)

type Login struct {
	UserName string        `json:"username" binding:"required"`
	Password string        `json:"password" binding:"required"`
	Tenant   models.Tenant `json:"-"`
}

type Auth struct {
	ClientID     string        `json:"clientId" binding:"required"`
	Scope        string        `json:"scope" binding:"required"`
	ResponseType string        `json:"responseType" binding:"required"`
	RedirectUri  string        `json:"redirectUri" binding:"required"`
	Nonce        string        `json:"nonce"`
	Tenant       models.Tenant `json:"-"`
	UserID       string        `json:"-"`
}

type Token struct {
	ClientID     string        `json:"clientId" binding:"required"`
	ClientSecret string        `json:"clientSecret" binding:"required"`
	GrantType    string        `json:"grantType" binding:"required"`
	Code         string        `json:"code" binding:"required"`
	RedirectUri  string        `json:"redirectUri" binding:"required"`
	State        string        `json:"state" binding:"required"`
	Nonce        string        `json:"nonce" binding:"required"`
	Scope        string        `json:"-"`
	UserID       string        `json:"-"`
	Tenant       models.Tenant `json:"-"`
	Client       models.Client `json:"-"`
}
