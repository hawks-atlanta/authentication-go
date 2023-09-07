package router

import (
	"net/http"
	"testing"

	"github.com/hawks-atlanta/authentication-go/database"
	"github.com/hawks-atlanta/authentication-go/internal/utils/jwt"
	"github.com/hawks-atlanta/authentication-go/models"
	"gorm.io/gorm"
)

func TestRouter_Authorization(t *testing.T) {
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

		expect.
			PATCH(ChallengeRoute).
			WithHeader(AuthorizationHeader, Bearer(tok)).
			Expect().
			Status(http.StatusOK)
	}))
	t.Run("No Token", database.Test(func(t *testing.T, db *gorm.DB) {
		expect, closeFunc := NewDefault(t, db)
		defer closeFunc()

		expect.
			PATCH(ChallengeRoute).
			WithHeader(AuthorizationHeader, "").
			Expect().
			Status(http.StatusUnauthorized)
	}))
	t.Run("Invalid Session", database.Test(func(t *testing.T, db *gorm.DB) {
		expect, closeFunc := NewDefault(t, db)
		defer closeFunc()

		u := models.RandomUser()
		expect.
			PATCH(ChallengeRoute).
			WithHeader(AuthorizationHeader, jwt.New([]byte("BYTE")).New(u.Claims())).
			Expect().
			Status(http.StatusUnauthorized)
	}))
}
