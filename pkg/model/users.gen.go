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
	ID            string    `gorm:"column:id;type:character(32);primaryKey" json:"id"`
	UserPoolID    int64     `gorm:"column:user_pool_id;type:bigint;not null" json:"userPoolId"`
	Username      string    `gorm:"column:username;type:character varying(127);not null" json:"username"`
	Password      string    `gorm:"column:password;type:character varying(127);not null" json:"password"`
	NickName      string    `gorm:"column:nick_name;type:character varying(127);not null" json:"nickName"`
	DisplayName   string    `gorm:"column:display_name;type:character varying(127);not null" json:"displayName"`
	Gender        string    `gorm:"column:gender;type:character(1);not null" json:"gender"`
	Birthdate     time.Time `gorm:"column:birthdate;type:date;not null" json:"birthdate"`
	Email         string    `gorm:"column:email;type:character varying(127);not null" json:"email"`
	EmailVerified bool      `gorm:"column:email_verified;type:boolean;not null" json:"emailVerified"`
	Phone         string    `gorm:"column:phone;type:character varying(20);not null" json:"phone"`
	PhoneVerified bool      `gorm:"column:phone_verified;type:boolean;not null" json:"phoneVerified"`
	Addr          string    `gorm:"column:addr;type:character varying(255);not null" json:"addr"`
	Avatar        string    `gorm:"column:avatar;type:character varying(255);not null" json:"avatar"`
	Type          int32     `gorm:"column:type;type:integer;not null" json:"type"`
	IsDisabled    bool      `gorm:"column:is_disabled;type:boolean;not null" json:"isDisabled"`
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now()" json:"updatedAt"`
	UserPool      UserPool  `json:"userPool"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
