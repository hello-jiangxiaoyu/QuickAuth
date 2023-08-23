// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"github.com/lib/pq"
)

const TableNameCode = "codes"

// Code mapped from table <codes>
type Code struct {
	ID        int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	UserID    string         `gorm:"column:user_id;type:character(32);not null" json:"userId"`
	AppID     string         `gorm:"column:app_id;type:character(32);not null" json:"appId"`
	Code      string         `gorm:"column:code;type:character(32);not null" json:"code"`
	Scope     pq.StringArray `gorm:"column:scope;type:character varying(255);not null" json:"scope"`
	State     string         `gorm:"column:state;type:character(63);not null" json:"state"`
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now()" json:"updatedAt"`
}

// TableName Code's table name
func (*Code) TableName() string {
	return TableNameCode
}
