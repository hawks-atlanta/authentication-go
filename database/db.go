package database

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	SQLite     = "sqlite"
	PostgreSQL = "postgres"
)

func NewDB(engine, dsn string) (db *gorm.DB, err error) {
	var dialector gorm.Dialector
	switch engine {
	case SQLite:
		dialector = sqlite.Open(dsn)
	case PostgreSQL:
		dialector = postgres.Open(dsn)
	default:
		err = fmt.Errorf("invalid engine `%s`, available are: postgres and sqlite", engine)
		return db, err
	}

	db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		err = fmt.Errorf("failed to initialize gorm: %w", err)
	}
	return db, err
}
