package request

import (
	"QuickAuth/pkg/model"
)

type (
	Login struct {
		UserName string       `form:"username" binding:"required"`
		Password string       `form:"password" binding:"required"`
		Tenant   model.Tenant `json:"-"`
	}

	// LoginProvider 第三方登录
	LoginProvider struct {
		ProviderName string       `uri:"provider"`
		Code         string       `query:"code"`
		Tenant       model.Tenant `query:"tenant"`
	}

	Auth struct {
		ClientID     string       `json:"client_id" binding:"required"`
		Scope        string       `json:"scope" binding:"required"`
		ResponseType string       `json:"response_type" binding:"required"`
		RedirectUri  string       `json:"redirect_uri" binding:"required"`
		Nonce        string       `json:"nonce"`
		Tenant       model.Tenant `json:"-"`
		UserID       string       `json:"-"`
	}

	Token struct {
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
		App          model.App    `json:"-"`
	}

	ProviderReq struct {
		Tenant       model.Tenant `json:"-"`
		ProviderId   int64        `json:"-" uri:"providerId"`
		Type         string       `json:"type"`
		AgentID      string       `json:"agentId"`
		ClientID     string       `json:"clientId"`
		ClientSecret string       `json:"clientSecret"`
	}
)

func (p *ProviderReq) ToModel() *model.Provider {
	return &model.Provider{
		TenantID:     p.Tenant.ID,
		Type:         p.Type,
		AppID:        p.Tenant.AppID,
		AgentID:      p.AgentID,
		ClientID:     p.ClientID,
		ClientSecret: p.ClientSecret,
	}
}
