package request

import (
	"QuickAuth/pkg/model"
)

type Login struct {
	UserName string       `json:"username" binding:"required"`
	Password string       `json:"password" binding:"required"`
	Tenant   model.Tenant `json:"-"`
}

// LoginProvider 第三方登录
type LoginProvider struct {
	ProviderName string       `uri:"provider"`
	Code         string       `query:"code"`
	Tenant       model.Tenant `query:"tenant"`
}

type Auth struct {
	ClientID     string       `json:"clientId" binding:"required"`
	Scope        string       `json:"scope" binding:"required"`
	ResponseType string       `json:"responseType" binding:"required"`
	RedirectUri  string       `json:"redirectUri" binding:"required"`
	Nonce        string       `json:"nonce"`
	Tenant       model.Tenant `json:"-"`
	UserID       string       `json:"-"`
}

type Token struct {
	ClientID     string       `json:"clientId" binding:"required"`
	ClientSecret string       `json:"clientSecret" binding:"required"`
	GrantType    string       `json:"grantType" binding:"required"`
	Code         string       `json:"code" binding:"required"`
	RedirectUri  string       `json:"redirectUri" binding:"required"`
	State        string       `json:"state" binding:"required"`
	Nonce        string       `json:"nonce" binding:"required"`
	Scope        string       `json:"-"`
	UserID       string       `json:"-"`
	Tenant       model.Tenant `json:"-"`
	Client       model.Client `json:"-"`
}
