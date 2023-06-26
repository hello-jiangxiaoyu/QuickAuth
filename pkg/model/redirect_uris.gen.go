// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameRedirectURI = "redirect_uris"

// RedirectURI mapped from table <redirect_uris>
type RedirectURI struct {
	ID         int32     `gorm:"column:id;type:integer;primaryKey;autoIncrement:true" json:"id"`
	ClientID   string    `gorm:"column:client_id;type:uuid;not null" json:"clientId"`
	URI        string    `gorm:"column:uri;type:character(63);not null" json:"uri"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp with time zone;not null;default:now()" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp with time zone;not null;default:now()" json:"updateTime"`
}

// TableName RedirectURI's table name
func (*RedirectURI) TableName() string {
	return TableNameRedirectURI
}