// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameClientSecret = "client_secrets"

// ClientSecret mapped from table <client_secrets>
type ClientSecret struct {
	ID         int32      `gorm:"column:id;type:integer;primaryKey;autoIncrement:true" json:"id"`
	ClientID   *string    `gorm:"column:client_id;type:uuid" json:"clientId"`
	Secret     *string    `gorm:"column:secret;type:character(63)" json:"secret"`
	Describe   *string    `gorm:"column:describe;type:character varying(127)" json:"describe"`
	CreateTime *time.Time `gorm:"column:create_time;type:timestamp with time zone;default:now()" json:"createTime"`
	UpdateTime *time.Time `gorm:"column:update_time;type:timestamp with time zone;default:now()" json:"updateTime"`
}

// TableName ClientSecret's table name
func (*ClientSecret) TableName() string {
	return TableNameClientSecret
}
