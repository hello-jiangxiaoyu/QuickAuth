package global

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true, // 关闭默认事务
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
