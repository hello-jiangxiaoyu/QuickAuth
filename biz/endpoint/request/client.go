package request

import (
	"QuickAuth/biz/endpoint/model"
	"github.com/lib/pq"
)

type AppReq struct {
	AppId    string `json:"-" form:"-" uri:"appId"`
	Name     string `json:"name" binding:"required"`
	Tag      string `json:"tag" binding:"required"`
	Icon     string `json:"icon" binding:"required"`
	Describe string `json:"describe" binding:"required"`
	Host     string `json:"host" binding:"required"`
	PoolId   int64  `json:"poolId"`
}

func (c *AppReq) ToModel() *model.App {
	return &model.App{
		Name:     c.Name,
		Tag:      c.Tag,
		Icon:     c.Icon,
		Describe: c.Describe,
	}
}

type AppSecretReq struct {
	AppId         string   `json:"-" uri:"appId"`
	SecretId      int64    `json:"-" uri:"secretId"`
	Describe      string   `json:"describe"`
	AccessExpire  int32    `json:"accessExpire"`
	RefreshExpire int32    `json:"refreshExpire"`
	Scope         []string `json:"scope"`
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
	AppID         string         `json:"-" form:"-" uri:"appId" binding:"required"`
	TenantID      int64          `json:"-" form:"-" uri:"tenantId"`
	UserPoolID    int64          `json:"userPoolId" binding:"required"`
	Type          int32          `json:"type"`
	Name          string         `json:"name" binding:"required"`
	Host          string         `json:"host" binding:"required"`
	Company       string         `json:"company" binding:"required"`
	GrantTypes    []string       `json:"grantTypes"`
	RedirectUris  pq.StringArray `json:"-"`
	CodeExpire    int32          `json:"codeExpire"`
	IDExpire      int32          `json:"idExpire"`
	AccessExpire  int32          `json:"accessExpire"`
	RefreshExpire int32          `json:"refreshExpire"`
	IsCode        bool           `json:"isCode"`
	IsRefresh     bool           `json:"isRefresh"`
	IsPassword    bool           `json:"isPassword"`
	IsCredential  bool           `son:"isCredential"`
	IsDeviceFlow  bool           `json:"isDeviceFlow"`
	Describe      string         `json:"describe"`
	IsDisabled    bool           `json:"isDisabled"`
}

func (t *TenantReq) ToModel() *model.Tenant {
	return &model.Tenant{
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
		IsDisabled:    t.IsDisabled,
	}
}
