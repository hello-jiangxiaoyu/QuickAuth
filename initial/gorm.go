package initial

import (
	"QuickAuth/conf"
	"QuickAuth/global"
	"QuickAuth/server/model"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

func MigrateDatabase() error {
	if !conf.IsFirstDeploy() {
		return nil
	}

	migrateList := []any{
		model.User{},
		model.UserPool{},
		model.App{},
		model.Tenant{},
	}

	if global.DB == nil {
		return errors.New("global db is nil, failed to migrate database")
	}

	if err := global.DB.AutoMigrate(migrateList...); err != nil {
		fmt.Println("migrate db err: ", err)
		return err
	}

	return nil
}

func NewGormDB(dbType string, dsn string) (*gorm.DB, error) {
	switch dbType {
	case conf.DBPostgres, conf.DBTimeScale:
		return getGormDB(postgres.Open(dsn))
	case conf.DBMySQL:
		return getGormDB(mysql.Open(dsn))
	case conf.DBSqlite:
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
