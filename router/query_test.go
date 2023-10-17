package router

import (
	"net/http"
	"testing"

	"github.com/hawks-atlanta/authentication-go/controller"
	"github.com/hawks-atlanta/authentication-go/database"
	"github.com/hawks-atlanta/authentication-go/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestRouter_UserByUsername(t *testing.T) {
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
			GET(UserUUIDRoute+"/"+*u.Username).
			WithHeader(AuthorizationHeader, Bearer(tok)).
			Expect().
			Status(http.StatusOK)
	}))
	t.Run("Invalid USERNAME", database.Test(func(t *testing.T, db *gorm.DB) {
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
			GET(UserUUIDRoute+"/INVALID").
			WithHeader(AuthorizationHeader, Bearer(tok)).
			Expect().
			Status(http.StatusInternalServerError)
	}))
}

func TestRouter_UsernameByUUIDRoute(t *testing.T) {
	t.Run("Succeed", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)

		expect, closeFunc := NewDefault(t, db)
		defer closeFunc()

		u := models.RandomUser()
		c, err := controller.New(controller.WithDB(db))
		assertions.Nil(err)

		err = c.Register(u)
		assertions.Nil(err)

		tok := expect.
			POST(LoginRoute).
			WithJSON(u).
			Expect().
			Status(http.StatusCreated).
			JSON().
			Object().
			Value("jwt").
			String().
			Raw()

		expect.
			GET(UsernameByUUIDRoute+"/"+u.UUID.String()).
			WithHeader(AuthorizationHeader, Bearer(tok)).
			Expect().
			Status(http.StatusOK).
			JSON().Object().Value("username").IsEqual(u.Username)
	}))
	t.Run("Invalid USERNAME", database.Test(func(t *testing.T, db *gorm.DB) {
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
			GET(UsernameByUUIDRoute+"/INVALID").
			WithHeader(AuthorizationHeader, Bearer(tok)).
			Expect().
			Status(http.StatusInternalServerError)
	}))
}
