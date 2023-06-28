package orm

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

const (
	DBPostgres  = "postgres"
	DBTimeScale = "timescale"
	DBMySQL     = "mysql"
	DBSqlite    = "sqlite"
)

func NewGormDB(dbType string, dsn string) (*gorm.DB, error) {
	switch dbType {
	case DBPostgres, DBTimeScale:
		return getGormDB(postgres.Open(dsn))
	case DBMySQL:
		return getGormDB(mysql.Open(dsn))
	case DBSqlite:
		return getGormDB(sqlite.Open(dsn))
	default:
		return nil, errors.New("unsupported database type")
	}
}

func getGormDB(d gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(d, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	return gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
}
