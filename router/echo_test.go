package router

import (
	"net/http"
	"testing"

	"github.com/hawks-atlanta/authentication-go/database"
	"gorm.io/gorm"
)

func TestRouter_AnyEcho(t *testing.T) {
	t.Run("GET", database.Test(func(t *testing.T, db *gorm.DB) {
		expect, closeFunc := NewDefault(t, db)
		defer closeFunc()

		expect.
			GET(EchoRoute).
			Expect().
			Status(http.StatusOK)
	}))
	t.Run("POST", database.Test(func(t *testing.T, db *gorm.DB) {
		expect, closeFunc := NewDefault(t, db)
		defer closeFunc()

		expect.
			POST(EchoRoute).
			Expect().
			Status(http.StatusOK)
	}))
}
