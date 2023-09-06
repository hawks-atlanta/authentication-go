package database

import (
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var nextDB atomic.Int64

func nextSQLite() string {
	return fmt.Sprintf("file:memdb%d?mode=memory&cache=shared", nextDB.Add(1))
}

const postgresDSN = "host=127.0.0.1 user=username password=password dbname=database port=5432 sslmode=disable"

func Close(t *testing.T, db *gorm.DB) {
	conn, err := db.DB()
	assert.Nil(t, err)
	conn.Close()
}

func Test(testFunc func(*testing.T, *gorm.DB)) func(*testing.T) {
	return func(t *testing.T) {
		t.Run(SQLite, func(t *testing.T) {
			db, err := NewDB(SQLite, nextSQLite())
			assert.Nil(t, err)
			defer Close(t, db)

			testFunc(t, db)
		})
		t.Run(PostgreSQL, func(t *testing.T) {
			db, err := NewDB(PostgreSQL, postgresDSN)
			assert.Nil(t, err)
			defer Close(t, db)

			testFunc(t, db)
		})
	}
}
