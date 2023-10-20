package router

import (
	"net/http"
	"testing"

	"github.com/hawks-atlanta/authentication-go/database"
	"github.com/hawks-atlanta/authentication-go/models"
	"gorm.io/gorm"
)

func TestRouter_Register(t *testing.T) {
	t.Run("Succeed", database.Test(func(t *testing.T, db *gorm.DB) {
		expect, closeFunc := NewDefault(t, db)
		defer closeFunc()

		u := models.RandomUser()
		expect.
			POST(RegisterRoute).
			WithJSON(u).
			Expect().
			Status(http.StatusCreated).
			JSON().
			Object().
			Value("jwt").
			String().
			NotEmpty()
	}))
	t.Run("Repeated user", database.Test(func(t *testing.T, db *gorm.DB) {
		expect, closeFunc := NewDefault(t, db)
		defer closeFunc()

		u := models.RandomUser()
		expect.
			POST(RegisterRoute).
			WithJSON(u).
			Expect().
			Status(http.StatusCreated).
			JSON().
			Object().
			Value("jwt").
			String().
			NotEmpty()
		expect.
			POST(RegisterRoute).
			WithJSON(u).
			Expect().
			Status(http.StatusConflict)
	}))
	t.Run("Invalid JSON", database.Test(func(t *testing.T, db *gorm.DB) {
		expect, closeFunc := NewDefault(t, db)
		defer closeFunc()

		expect.
			POST(RegisterRoute).
			WithText("}").
			WithHeader("Content-Type", "application/json").
			Expect().
			Status(http.StatusBadRequest)
	}))
}
