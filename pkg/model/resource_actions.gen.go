// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameResourceAction = "resource_actions"

// ResourceAction mapped from table <resource_actions>
type ResourceAction struct {
	ID        int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	TenantID  int64     `gorm:"column:tenant_id;type:bigint;not null" json:"tenantId"`
	AppID     string    `gorm:"column:app_id;type:character(32);not null" json:"appId"`
	Type      string    `gorm:"column:type;type:character varying(32);not null" json:"type"`
	Code      string    `gorm:"column:code;type:character varying(63);not null" json:"code"`
	Name      string    `gorm:"column:name;type:character varying(127);not null" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now()" json:"updatedAt"`
	App       App       `json:"app"`
	Tenant    Tenant    `json:"tenant"`
}

// TableName ResourceAction's table name
func (*ResourceAction) TableName() string {
	return TableNameResourceAction
}
