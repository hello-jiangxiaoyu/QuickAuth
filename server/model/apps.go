package model

import (
	"time"
)

// App mapped from table <apps>
type App struct {
	ID         string     `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	Name       *string    `gorm:"column:name;type:character varying(255)" json:"name"`
	Describe   *string    `gorm:"column:describe;type:character varying(255)" json:"describe"`
	CreateTime *time.Time `gorm:"column:create_time;type:timestamp with time zone;default:now()" json:"createTime"`
	UpdateTime *time.Time `gorm:"column:update_time;type:timestamp with time zone;default:now()" json:"updateTime"`
}

// TableName App's table name
func (*App) TableName() string {
	return "apps"
}
