package orm

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
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

type Writer struct {
	logger.Writer
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
	// todo: user zap log gorm sql
	_default := logger.New(&Writer{Writer: log.New(os.Stdout, "\r\n", log.LstdFlags)}, logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})

	return gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: false,
		Logger:                 _default.LogMode(logger.Info),
	})
}
