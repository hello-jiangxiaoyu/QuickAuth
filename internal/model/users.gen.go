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
	ID          string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserPoolID  string    `gorm:"column:user_pool_id;type:uuid;not null" json:"userPoolId"`
	Username    string    `gorm:"column:username;type:character varying(127);not null" json:"username"`
	Password    *string   `gorm:"column:password;type:character varying(127)" json:"password"`
	DisplayName *string   `gorm:"column:display_name;type:character varying(127)" json:"displayName"`
	Email       *string   `gorm:"column:email;type:character varying(127)" json:"email"`
	Phone       *string   `gorm:"column:phone;type:character varying(20)" json:"phone"`
	IsDisabled  bool      `gorm:"column:is_disabled;type:boolean;not null" json:"isDisabled"`
	CreateTime  time.Time `gorm:"column:create_time;type:timestamp with time zone;not null;default:now()" json:"createTime"`
	UpdateTime  time.Time `gorm:"column:update_time;type:timestamp with time zone;not null;default:now()" json:"updateTime"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
