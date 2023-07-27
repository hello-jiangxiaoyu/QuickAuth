package request

import (
	"QuickAuth/pkg/model"
	"github.com/lib/pq"
)

type AppReq struct {
	AppId    string `json:"-" form:"-" uri:"appId"`
	Name     string `json:"name"`
	Tag      string `json:"tag"`
	Icon     string `json:"icon"`
	Describe string `json:"describe"`
}

func (c *AppReq) ToModel() model.App {
	return model.App{
		Name:     c.Name,
		Tag:      c.Tag,
		Icon:     c.Icon,
		Describe: c.Describe,
	}
}

type AppSecretReq struct {
	AppId         string `json:"-" uri:"appId"`
	SecretId      int64  `json:"-" uri:"secretId"`
	Describe      string `json:"describe"`
	AccessExpire  int32  `json:"accessExpire"`
	RefreshExpire int32  `json:"refreshExpire"`
	Scope         string `json:"scope"`
}

func (c *AppSecretReq) ToModel() model.AppSecret {
	return model.AppSecret{
		Describe:      c.Describe,
		Scope:         c.Scope,
		AccessExpire:  c.AccessExpire,
		RefreshExpire: c.RefreshExpire,
	}
}

type RedirectUriReq struct {
	Tenant model.Tenant `json:"-"`
	UriId  uint         `json:"-" uri:"uriId"`
	Uri    string       `json:"uri"`
}

type TenantReq struct {
	AppID         string         `json:"-" form:"-" uri:"appId"`
	TenantID      int64          `json:"-" form:"-" uri:"tenantId"`
	UserPoolID    int64          `json:"userPoolId"`
	Type          int32          `json:"type"`
	Name          string         `json:"name"`
	Host          string         `json:"host"`
	Company       string         `json:"company"`
	GrantTypes    []string       `json:"grantTypes"`
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
	Describe      string         `json:"describe"`
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
