package router

import (
	"net/http"
	"testing"
	"time"

	"github.com/hawks-atlanta/authentication-go/controller"
	"github.com/hawks-atlanta/authentication-go/database"
	"github.com/hawks-atlanta/authentication-go/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestLogs(t *testing.T) {
	t.Run("Succeed", database.Test(func(t *testing.T, db *gorm.DB) {

		expect, closeFunc := NewDefault(t, db)
		defer closeFunc()

		expect.
			GET(LogsRoute).
			Expect().
			Status(http.StatusOK).
			JSON().
			Array()
	}))
}

func TestUserLogs(t *testing.T) {

	t.Run("Succeed", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)

		expect, closeFunc := NewDefault(t, db)
		defer closeFunc()

		l, u := models.RandomLog()
		assertions.Nil(db.Create(l).Error)

		expect.
			POST(LogsUserRoute).
			WithJSON(controller.Filter[models.User]{Object: models.User{Username: &u}, ItemsPerPage: 10, Page: 1}).
			Expect().
			Status(http.StatusOK).
			JSON().
			Array()
	}))

	t.Run("Logs not found", database.Test(func(t *testing.T, db *gorm.DB) {

		expect, closeFunc := NewDefault(t, db)
		defer closeFunc()

		u := "Test"
		expect.
			POST(LogsUserRoute).
			WithJSON(controller.Filter[models.User]{Object: models.User{Username: &u}, ItemsPerPage: 10, Page: 1}).
			Expect().
			Status(http.StatusNotFound).
			JSON().
			Object().
			Value("succeed").
			Boolean().
			IsFalse()
	}))
	t.Run("Invalid JSON", database.Test(func(t *testing.T, db *gorm.DB) {
		expect, closeFunc := NewDefault(t, db)
		defer closeFunc()

		expect.
			POST(LogsUserRoute).
			WithText("}").
			WithHeader("Content-Type", "application/json").
			Expect().
			Status(http.StatusBadRequest)
	}))
}

func TestDateLogs(t *testing.T) {

	t.Run("Succeed", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)

		expect, closeFunc := NewDefault(t, db)
		defer closeFunc()

		l, _ := models.RandomLog()
		assertions.Nil(db.Create(l).Error)

		expect.
			POST(LogsDateRoute).
			WithJSON(controller.Filter[time.Time]{Object: time.Now().Add(-time.Hour * 1), ItemsPerPage: 10, Page: 1}).
			Expect().
			Status(http.StatusOK)
	}))

	t.Run("Invalid JSON", database.Test(func(t *testing.T, db *gorm.DB) {
		expect, closeFunc := NewDefault(t, db)
		defer closeFunc()

		expect.
			POST(LogsDateRoute).
			WithText("}").
			WithHeader("Content-Type", "application/json").
			Expect().
			Status(http.StatusBadRequest)
	}))
}
