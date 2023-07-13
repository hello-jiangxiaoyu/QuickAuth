package request

import (
	"QuickAuth/pkg/model"
	"github.com/lib/pq"
)

type AppReq struct {
	AppId         string         `json:"-" uri:"appId"`
	Name          string         `json:"name"`
	Describe      string         `json:"describe"`
	GrantTypes    pq.StringArray `json:"grantTypes"`
	RedirectUris  pq.StringArray `json:"redirectUris"`
	TokenExpire   int32          `json:"tokenExpire"`
	RefreshExpire int32          `json:"refreshExpire"`
	CodeExpire    int32          `json:"codeExpire"`
}

func (c *AppReq) ToModel() model.App {
	return model.App{
		ID:            c.AppId,
		Name:          c.Name,
		Describe:      &c.Describe,
		GrantTypes:    c.GrantTypes,
		RedirectUris:  c.RedirectUris,
		TokenExpire:   c.TokenExpire,
		RefreshExpire: c.RefreshExpire,
		CodeExpire:    c.CodeExpire,
	}
}

type AppSecretReq struct {
	AppId    string `json:"-" uri:"appId"`
	Describe string `json:"describe"`
}

func (c *AppSecretReq) ToModel() model.AppSecret {
	return model.AppSecret{
		AppID:    c.AppId,
		Describe: &c.Describe,
	}
}

type RedirectUriReq struct {
	AppId    string `json:"-" uri:"appId"`
	TenantId string `json:"-" uri:"tenantId"`
	UriId    uint   `json:"-" uri:"uriId"`
	Uri      string `json:"uri"`
}

type TenantReq struct {
	TenantID   int64   `json:"-" uri:"tenantId"`
	AppId      string  `json:"-" uri:"appId"`
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
		AppId:      t.AppId,
		UserPoolID: t.UserPoolID,
		Type:       t.Type,
		Name:       t.Name,
		Host:       t.Host,
		Company:    t.Company,
		Describe:   t.Describe,
	}
}
