package model

import (
	"time"
)

// UserPool mapped from table <user_pools>
type UserPool struct {
	ID         string     `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	Name       *string    `gorm:"column:name;type:character varying(255)" json:"name"`
	Describe   *string    `gorm:"column:describe;type:character varying(255)" json:"describe"`
	CreateTime *time.Time `gorm:"column:create_time;type:timestamp with time zone;default:now()" json:"createTime"`
	UpdateTime *time.Time `gorm:"column:update_time;type:timestamp with time zone;default:now()" json:"updateTime"`
}

// TableName UserPool's table name
func (*UserPool) TableName() string {
	return "user_pools"
}
