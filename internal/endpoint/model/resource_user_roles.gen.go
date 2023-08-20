// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameResourceUserRole = "resource_user_roles"

// ResourceUserRole mapped from table <resource_user_roles>
type ResourceUserRole struct {
	ID         int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	TenantID   int64     `gorm:"column:tenant_id;type:bigint;not null" json:"tenantId"`
	ResourceID int64     `gorm:"column:resource_id;type:bigint;not null" json:"resourceId"`
	NodeID     int64     `gorm:"column:node_id;type:bigint;not null" json:"nodeId"`
	UserID     string    `gorm:"column:user_id;type:character(32);not null" json:"userId"`
	RoleID     int64     `gorm:"column:role_id;type:bigint;not null" json:"roleId"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now()" json:"updatedAt"`
	Tenant     Tenant    `json:"tenant"`
}

// TableName ResourceUserRole's table name
func (*ResourceUserRole) TableName() string {
	return TableNameResourceUserRole
}
