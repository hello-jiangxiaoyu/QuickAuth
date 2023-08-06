// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameProvider = "providers"

// Provider mapped from table <providers>
type Provider struct {
	ID           int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	TenantID     int64     `gorm:"column:tenant_id;type:bigint;not null" json:"tenantId"`
	AppID        string    `gorm:"column:app_id;type:character(32);not null" json:"appId"`
	Type         string    `gorm:"column:type;type:character varying(32);not null" json:"type"`
	ClientID     string    `gorm:"column:client_id;type:character varying(255);not null" json:"clientId"`
	ClientSecret string    `gorm:"column:client_secret;type:character varying(255);not null" json:"clientSecret"`
	AgentID      string    `gorm:"column:agent_id;type:character varying(255);not null" json:"agentId"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now()" json:"updatedAt"`
	App          App       `json:"app"`
	Tenant       Tenant    `json:"tenant"`
}

// TableName Provider's table name
func (*Provider) TableName() string {
	return TableNameProvider
}
