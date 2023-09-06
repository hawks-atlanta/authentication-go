package controller

import (
	"testing"

	"github.com/hawks-atlanta/authentication-go/database"
	"github.com/hawks-atlanta/authentication-go/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestController_Login(t *testing.T) {
	t.Run("Succeed", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)
		c, err := New(WithDB(db))
		assertions.Nil(err)

		u := models.RandomUser()
		err = db.
			Create(u).
			Error
		assertions.Nil(err)

		u2, err := c.Login(u)
		assertions.Nil(err)

		assertions.Equal(u.UUID, u2.UUID)
		assertions.Equal(*u.Username, *u2.Username)
		assertions.Equal(u.PasswordHash, u2.PasswordHash)
	}))
	t.Run("Invalid Username", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)
		c, err := New(WithDB(db))
		assertions.Nil(err)

		u := models.RandomUser()
		err = db.
			Create(u).
			Error
		assertions.Nil(err)

		wrong := "WRONG"
		_, err = c.Login(&models.User{Username: &wrong, Password: u.Password})
		assertions.NotNil(err)
	}))
	t.Run("Invalid Password", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)
		c, err := New(WithDB(db))
		assertions.Nil(err)

		u := models.RandomUser()
		err = db.
			Create(u).
			Error
		assertions.Nil(err)

		_, err = c.Login(&models.User{Username: u.Username, Password: "WRONG"})
		assertions.NotNil(err)
	}))
	t.Run("No Username", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)
		c, err := New(WithDB(db))
		assertions.Nil(err)

		u := models.RandomUser()
		err = db.
			Create(u).
			Error
		assertions.Nil(err)

		_, err = c.Login(&models.User{Password: u.Password})
		assertions.NotNil(err)
	}))
}
