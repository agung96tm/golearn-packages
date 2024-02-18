package lib

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
)

type Database struct {
	ORM *gorm.DB
}

type DatabaseConfig struct {
	DSN string
}

func getDSN(dsn string) string {
	if strings.HasPrefix(dsn, "sqlite3://") {
		return strings.TrimPrefix(dsn, "sqlite3://")
	}
	return dsn
}

func NewDB(config DatabaseConfig) (*Database, error) {
	db, err := gorm.Open(sqlite.Open(getDSN(config.DSN)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	d, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := d.Ping(); err != nil {
		return nil, err
	}

	return &Database{
		ORM: db,
	}, nil
}
