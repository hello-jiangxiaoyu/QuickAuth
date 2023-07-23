// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID          string    `gorm:"column:id;type:character(32);primaryKey" json:"id"`
	UserPoolID  int64     `gorm:"column:user_pool_id;type:bigint;not null" json:"userPoolId"`
	Username    string    `gorm:"column:username;type:character varying(127);not null" json:"username"`
	Password    string    `gorm:"column:password;type:character varying(127);not null" json:"password"`
	DisplayName string    `gorm:"column:display_name;type:character varying(127);not null" json:"displayName"`
	Email       string    `gorm:"column:email;type:character varying(127);not null" json:"email"`
	Phone       string    `gorm:"column:phone;type:character varying(20);not null" json:"phone"`
	Type        int32     `gorm:"column:type;type:integer;not null" json:"type"`
	IsDisabled  int32     `gorm:"column:is_disabled;type:integer;not null" json:"isDisabled"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now()" json:"updatedAt"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
