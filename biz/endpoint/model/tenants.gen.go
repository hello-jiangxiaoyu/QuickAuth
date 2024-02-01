// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"github.com/lib/pq"
)

const TableNameTenant = "tenants"

// Tenant mapped from table <tenants>
type Tenant struct {
	ID            int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	AppID         string         `gorm:"column:app_id;type:character(32);not null" json:"appId"`
	UserPoolID    int64          `gorm:"column:user_pool_id;type:bigint;not null" json:"userPoolId"`
	Type          int32          `gorm:"column:type;type:integer;not null" json:"type"`
	Name          string         `gorm:"column:name;type:character varying(127);not null" json:"name"`
	Host          string         `gorm:"column:host;type:character varying(127);not null" json:"host"`
	Company       string         `gorm:"column:company;type:character varying(127);not null" json:"company"`
	GrantTypes    pq.StringArray `gorm:"column:grant_types;type:character varying(127)[];not null" json:"grantTypes"`
	RedirectUris  pq.StringArray `gorm:"column:redirect_uris;type:character varying(127)[];not null" json:"redirectUris"`
	CodeExpire    int32          `gorm:"column:code_expire;type:integer;not null;default:120" json:"codeExpire"`
	IDExpire      int32          `gorm:"column:id_expire;type:integer;not null;default:604800" json:"idExpire"`
	AccessExpire  int32          `gorm:"column:access_expire;type:integer;not null;default:604800" json:"accessExpire"`
	RefreshExpire int32          `gorm:"column:refresh_expire;type:integer;not null;default:2592000" json:"refreshExpire"`
	IsCode        bool           `gorm:"column:is_code;type:boolean;not null;default:true" json:"isCode"`
	IsRefresh     bool           `gorm:"column:is_refresh;type:boolean;not null;default:true" json:"isRefresh"`
	IsPassword    bool           `gorm:"column:is_password;type:boolean;not null" json:"isPassword"`
	IsCredential  bool           `gorm:"column:is_credential;type:boolean;not null;default:true" json:"isCredential"`
	IsDeviceFlow  bool           `gorm:"column:is_device_flow;type:boolean;not null" json:"isDeviceFlow"`
	Config        string         `gorm:"column:config;type:jsonb;not null" json:"config"`
	Describe      string         `gorm:"column:describe;type:character varying(127);not null" json:"describe"`
	IsDisabled    bool           `gorm:"column:is_disabled;type:boolean;not null" json:"isDisabled"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"createdAt"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now()" json:"updatedAt"`
	App           App            `json:"app"`
	UserPool      UserPool       `json:"userPool"`
}

// TableName Tenant's table name
func (*Tenant) TableName() string {
	return TableNameTenant
}