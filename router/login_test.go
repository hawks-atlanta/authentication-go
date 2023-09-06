package router

import (
	"net/http"
	"testing"

	"github.com/hawks-atlanta/authentication-go/database"
	"github.com/hawks-atlanta/authentication-go/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestRouter_Login(t *testing.T) {
	t.Run("Succeed", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)

		expect, closeFunc := NewDefault(t, db)
		defer closeFunc()

		u := models.RandomUser()
		assertions.Nil(db.Create(u).Error)

		expect.
			POST(LoginRoute).
			WithJSON(models.User{Username: u.Username, Password: u.Password}).
			Expect().
			Status(http.StatusCreated).
			JSON().
			Object().
			Value("jwt").
			String().
			NotEmpty()
	}))
	t.Run("Invalid credentials", database.Test(func(t *testing.T, db *gorm.DB) {
		expect, closeFunc := NewDefault(t, db)
		defer closeFunc()

		wrong := "wrong"
		expect.
			POST(LoginRoute).
			WithJSON(models.User{Username: &wrong, Password: "INVALID"}).
			Expect().
			Status(http.StatusUnauthorized).
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
			POST(LoginRoute).
			WithText("}").
			WithHeader("Content-Type", "application/json").
			Expect().
			Status(http.StatusBadRequest)
	}))
}
