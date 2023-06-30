package request

import (
	"QuickAuth/pkg/model"
	"github.com/lib/pq"
)

type ClientReq struct {
	ClientId      string         `json:"-" uri:"clientId"`
	Name          string         `json:"name"`
	Describe      string         `json:"describe"`
	GrantTypes    pq.StringArray `json:"grantTypes"`
	RedirectUris  pq.StringArray `json:"redirectUris"`
	TokenExpire   int32          `json:"tokenExpire"`
	RefreshExpire int32          `json:"refreshExpire"`
	CodeExpire    int32          `json:"codeExpire"`
}

func (c *ClientReq) ToModel() model.Client {
	return model.Client{
		ID:            c.ClientId,
		Name:          c.Name,
		Describe:      &c.Describe,
		GrantTypes:    c.GrantTypes,
		RedirectUris:  c.RedirectUris,
		TokenExpire:   c.TokenExpire,
		RefreshExpire: c.RefreshExpire,
		CodeExpire:    c.CodeExpire,
	}
}

type ClientSecretReq struct {
	ClientId string `json:"-" uri:"clientId"`
	Describe string `json:"describe"`
}

func (c *ClientSecretReq) ToModel() model.ClientSecret {
	return model.ClientSecret{
		ClientID: c.ClientId,
		Describe: &c.Describe,
	}
}

type RedirectUriReq struct {
	ClientId string `json:"-" uri:"clientId"`
	UriId    uint   `json:"-" uri:"uriId"`
	Uri      string `json:"uri"`
}

type TenantReq struct {
	TenantID   int64   `json:"-" uri:"tenantId"`
	ClientID   string  `json:"-" uri:"clientId"`
	UserPoolID int64   `json:"userPoolId"`
	Type       int32   `json:"type"`
	Name       string  `json:"name"`
	Host       string  `json:"host"`
	Company    string  `json:"company"`
	Describe   *string `json:"describe"`
}

func (t *TenantReq) ToModel() model.Tenant {
	return model.Tenant{
		ID:         t.TenantID,
		ClientID:   t.ClientID,
		UserPoolID: t.UserPoolID,
		Type:       t.Type,
		Name:       t.Name,
		Host:       t.Host,
		Company:    t.Company,
		Describe:   t.Describe,
	}
}
