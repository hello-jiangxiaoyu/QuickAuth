package model

import (
	"time"
)

type User struct {
	ID          string     `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserPoolID  string     `gorm:"column:user_pool_id;type:uuid;not null" json:"userPoolId"`
	Username    string     `gorm:"column:username;type:character varying(255);not null" json:"username"`
	Password    *string    `gorm:"column:password;type:character varying(255)" json:"password"`
	DisplayName *string    `gorm:"column:display_name;type:character varying(255)" json:"displayName"`
	Email       *string    `gorm:"column:email;type:character varying(255)" json:"email"`
	Phone       *string    `gorm:"column:phone;type:character varying(20)" json:"phone"`
	IsDisabled  *bool      `gorm:"column:is_disabled;type:boolean" json:"isDisabled"`
	CreateTime  *time.Time `gorm:"column:create_time;type:timestamp with time zone;default:now()" json:"createTime"`
	UpdateTime  *time.Time `gorm:"column:update_time;type:timestamp with time zone;default:now()" json:"updateTime"`
}

// TableName User's table name
func (*User) TableName() string {
	return "users"
}
