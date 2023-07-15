// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameApp = "apps"

// App mapped from table <apps>
type App struct {
	ID         string    `gorm:"column:id;type:character(32);primaryKey" json:"id"`
	Name       string    `gorm:"column:name;type:character varying(127);not null" json:"name"`
	Describe   string    `gorm:"column:describe;type:character varying(127);not null" json:"describe"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp with time zone;not null;default:now()" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp with time zone;not null;default:now()" json:"updateTime"`
}

// TableName App's table name
func (*App) TableName() string {
	return TableNameApp
}
