package router

import (
	"net/http"
	"testing"

	"github.com/hawks-atlanta/authentication-go/controller"
	"github.com/hawks-atlanta/authentication-go/database"
	"github.com/hawks-atlanta/authentication-go/internal/utils/random"
	"github.com/hawks-atlanta/authentication-go/models"
	"gorm.io/gorm"
)

func TestController_UpdatePassword(t *testing.T) {
	t.Run("Succeed", database.Test(func(t *testing.T, db *gorm.DB) {
		expect, closeFunc := NewDefault(t, db)
		defer closeFunc()

		u := models.RandomUser()
		tok := expect.
			POST(RegisterRoute).
			WithJSON(u).
			Expect().
			Status(http.StatusCreated).
			JSON().
			Object().
			Value("jwt").
			String().
			Raw()

		req := controller.UpdatePasswordRequest{
			OldPassword: u.Password,
			NewPassword: random.String(16),
		}
		expect.
			PATCH(AccountPasswordRoute).
			WithHeader(AuthorizationHeader, Bearer(tok)).
			WithJSON(req).
			Expect().
			Status(http.StatusOK)
	}))
	t.Run("Invalid Old Password", database.Test(func(t *testing.T, db *gorm.DB) {
		expect, closeFunc := NewDefault(t, db)
		defer closeFunc()

		u := models.RandomUser()
		tok := expect.
			POST(RegisterRoute).
			WithJSON(u).
			Expect().
			Status(http.StatusCreated).
			JSON().
			Object().
			Value("jwt").
			String().
			Raw()

		req := controller.UpdatePasswordRequest{
			OldPassword: "INVALID",
			NewPassword: random.String(16),
		}
		expect.
			PATCH(AccountPasswordRoute).
			WithHeader(AuthorizationHeader, Bearer(tok)).
			WithJSON(req).
			Expect().
			Status(http.StatusUnauthorized)
	}))
	t.Run("Invalid JWT", database.Test(func(t *testing.T, db *gorm.DB) {
		expect, closeFunc := NewDefault(t, db)
		defer closeFunc()

		u := models.RandomUser()
		tok := expect.
			POST(RegisterRoute).
			WithJSON(u).
			Expect().
			Status(http.StatusCreated).
			JSON().
			Object().
			Value("jwt").
			String().
			Raw()

		expect.
			PATCH(AccountPasswordRoute).
			WithHeader(AuthorizationHeader, Bearer(tok)).
			WithHeader("Content-Type", "application/json").
			WithText("}").
			Expect().
			Status(http.StatusBadRequest)
	}))
}
