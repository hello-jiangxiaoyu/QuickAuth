// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserPool = "user_pools"

// UserPool mapped from table <user_pools>
type UserPool struct {
	ID         int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	Name       string    `gorm:"column:name;type:character varying(127);not null" json:"name"`
	Describe   string    `gorm:"column:describe;type:character varying(127);not null" json:"describe"`
	IsDisabled int32     `gorm:"column:is_disabled;type:integer;not null" json:"isDisabled"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp with time zone;not null;default:now()" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp with time zone;not null;default:now()" json:"updateTime"`
}

// TableName UserPool's table name
func (*UserPool) TableName() string {
	return TableNameUserPool
}
