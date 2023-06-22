package model

import (
	"time"
)

// Tenant mapped from table <tenants>
type Tenant struct {
	ID         string     `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	AppID      string     `gorm:"column:app_id;type:uuid;not null" json:"appId"`
	UserPoolID string     `gorm:"column:user_pool_id;type:uuid;not null" json:"userPoolId"`
	Type       *int32     `gorm:"column:type;type:integer" json:"type"`
	Name       *string    `gorm:"column:name;type:character varying(255)" json:"name"`
	Host       string     `gorm:"column:host;type:character varying(255);not null" json:"host"`
	Company    string     `gorm:"column:company;type:character varying(255);not null" json:"company"`
	Describe   *string    `gorm:"column:describe;type:character varying(255)" json:"describe"`
	CreateTime *time.Time `gorm:"column:create_time;type:timestamp with time zone;default:now()" json:"createTime"`
	UpdateTime *time.Time `gorm:"column:update_time;type:timestamp with time zone;default:now()" json:"updateTime"`
}

// TableName Tenant's table name
func (*Tenant) TableName() string {
	return "tenants"
}
