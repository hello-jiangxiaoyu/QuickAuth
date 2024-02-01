// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"github.com/lib/pq"
)

const TableNameAppSecret = "app_secrets"

// AppSecret mapped from table <app_secrets>
type AppSecret struct {
	ID            int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	AppID         string         `gorm:"column:app_id;type:character(32);not null" json:"appId"`
	Secret        string         `gorm:"column:secret;type:character(63);not null" json:"secret"`
	Scope         pq.StringArray `gorm:"column:scope;type:character varying(127)[];not null" json:"scope"`
	AccessExpire  int32          `gorm:"column:access_expire;type:integer;not null;default:604800" json:"accessExpire"`
	RefreshExpire int32          `gorm:"column:refresh_expire;type:integer;not null;default:2592000" json:"refreshExpire"`
	Describe      string         `gorm:"column:describe;type:character varying(127);not null" json:"describe"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"createdAt"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now()" json:"updatedAt"`
}

// TableName AppSecret's table name
func (*AppSecret) TableName() string {
	return TableNameAppSecret
}