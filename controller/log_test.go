package controller

import (
	"testing"
	"time"

	"github.com/hawks-atlanta/authentication-go/database"
	"github.com/hawks-atlanta/authentication-go/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestLog(t *testing.T) {
	t.Run("Failed user logs query", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)
		c, err := New(WithDB(db))
		assertions.Nil(err)

		u := "Test"
		_, err = c.GetLogsByUser(&Filter[models.User]{Object: models.User{Username: &u}, ItemsPerPage: -10, Page: 1})
		assertions.EqualError(err, "failed to query user: record not found")
	}))

	t.Run("No available logs", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)
		c, err := New(WithDB(db))
		assertions.Nil(err)

		logs, err := c.GetLogs()
		assertions.Nil(err)
		assertions.Equal(len(logs), 0, "The logs length should be equal to zero")

		u := "Test"
		_, err = c.GetLogsByUser(&Filter[models.User]{Object: models.User{Username: &u}, ItemsPerPage: 10, Page: 1})
		assertions.NotNil(err)

		_, err = c.GetLogsByDate(&Filter[time.Time]{Object: time.Now(), ItemsPerPage: 10, Page: 1})
		assertions.Nil(err)
	}))

	t.Run("Succeed", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)
		c, err := New(WithDB(db))
		assertions.Nil(err)

		l, u := models.RandomLog()
		err = c.CreateLog(l)
		assertions.Nil(err)

		logs, err := c.GetLogs()
		assertions.Nil(err)

		assertions.Equal(len(logs), 1, "The logs length should be equal to 1")

		logs, err = c.GetLogsByUser(&Filter[models.User]{Object: models.User{Username: &u}, ItemsPerPage: 10, Page: 1})
		assertions.Nil(err)
		assertions.Equal(len(logs), 1, "The logs length filtered by user should be equal to 1")

		logs, err = c.GetLogsByDate(&Filter[time.Time]{Object: time.Now().Add(-time.Hour * 1), ItemsPerPage: 10, Page: 1})
		assertions.Nil(err)
		assertions.Equal(len(logs), 1, "The logs length filtered by user should be equal to 1")
	}))
}
