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
	Type         string    `gorm:"column:type;type:character varying(31);not null" json:"type"`
	ClientID     string    `gorm:"column:client_id;type:character varying(255);not null" json:"clientId"`
	ClientSecret string    `gorm:"column:client_secret;type:character varying(255);not null" json:"clientSecret"`
	AgentID      *string   `gorm:"column:agent_id;type:character varying(255)" json:"agentId"`
	CreateTime   time.Time `gorm:"column:create_time;type:timestamp with time zone;not null;default:now()" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:update_time;type:timestamp with time zone;not null;default:now()" json:"updateTime"`
}

// TableName Provider's table name
func (*Provider) TableName() string {
	return TableNameProvider
}
