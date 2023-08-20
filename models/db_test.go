package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewDB(t *testing.T) {
	t.Run(SQLite, func(t *testing.T) {
		assertions := assert.New(t)

		db, err := NewDB(SQLite, nextSQLite())
		assertions.Nil(err)
		Close(t, db)
	})
	t.Run(PostgreSQL, func(t *testing.T) {
		assertions := assert.New(t)

		db, err := NewDB(PostgreSQL, postgresDSN)
		assertions.Nil(err)
		Close(t, db)
	})
	t.Run("INVALID", func(t *testing.T) {
		assertions := assert.New(t)

		_, err := NewDB("INVALID", postgresDSN)
		assertions.NotNil(err)
	})
	t.Run("INVALID - DSN", func(t *testing.T) {
		assertions := assert.New(t)

		_, err := NewDB(PostgreSQL, "INVALID")
		assertions.NotNil(err)
	})
}
