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
		ID:       c.AppId,
		Name:     c.Name,
		Describe: &c.Describe,
	}
}

type AppSecretReq struct {
	AppId    string `json:"-" uri:"appId"`
	SecretId int64  `json:"-" uri:"secretId"`
	Describe string `json:"describe"`
	Scope    string `json:"scope"`
}

func (c *AppSecretReq) ToModel() model.AppSecret {
	return model.AppSecret{
		ID:       c.SecretId,
		AppID:    c.AppId,
		Describe: &c.Describe,
		Scope:    c.Scope,
	}
}

type RedirectUriReq struct {
	Tenant model.Tenant `json:"-"`
	AppId  string       `json:"-" uri:"appId"`
	UriId  uint         `json:"-" uri:"uriId"`
	Uri    string       `json:"uri"`
}

type TenantReq struct {
	Tenant        model.Tenant   `json:"-"`
	ID            int64          `json:"-"`
	AppID         string         `json:"-"`
	UserPoolID    int64          `json:"userPoolId"`
	Type          int32          `json:"type"`
	Name          string         `json:"name"`
	Host          string         `json:"host"`
	Company       string         `json:"company"`
	GrantTypes    pq.StringArray `json:"-"`
	RedirectUris  pq.StringArray `json:"-"`
	CodeExpire    int32          `json:"codeExpire"`
	IDExpire      int32          `json:"idExpire"`
	AccessExpire  int32          `json:"accessExpire"`
	RefreshExpire int32          `json:"refreshExpire"`
	IsCode        int32          `json:"isCode"`
	IsRefresh     int32          `json:"isRefresh"`
	IsPassword    int32          `json:"isPassword"`
	IsCredential  int32          `son:"isCredential"`
	IsDeviceFlow  int32          `json:"isDeviceFlow"`
	Describe      *string        `json:"describe"`
	IsDisabled    int32          `json:"isDisabled"`
}

func (t *TenantReq) ToModel() model.Tenant {
	return model.Tenant{
		AppID:         t.AppID,
		UserPoolID:    t.UserPoolID,
		Type:          t.Type,
		Name:          t.Name,
		Host:          t.Host,
		Company:       t.Company,
		GrantTypes:    t.GrantTypes,
		RedirectUris:  t.RedirectUris,
		CodeExpire:    t.CodeExpire,
		IDExpire:      t.IDExpire,
		AccessExpire:  t.AccessExpire,
		RefreshExpire: t.RefreshExpire,
		IsCode:        t.IsCode,
		IsRefresh:     t.IsRefresh,
		IsPassword:    t.IsPassword,
		IsCredential:  t.IsCredential,
		IsDeviceFlow:  t.IsDeviceFlow,
		Describe:      t.Describe,
	}
}
