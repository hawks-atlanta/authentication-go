package controller

import (
	"testing"

	"github.com/hawks-atlanta/authentication-go/database"
	"github.com/hawks-atlanta/authentication-go/internal/utils/random"
	"github.com/hawks-atlanta/authentication-go/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestController_UpdatePassword(t *testing.T) {
	t.Run("Succeed", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)

		c, err := New(WithDB(db), WithSecret(random.Bytes(16)))
		assertions.Nil(err)

		u := models.RandomUser()
		err = c.Register(u)
		assertions.Nil(err)

		req := UpdatePasswordRequest{
			OldPassword: u.Password,
			NewPassword: random.String(16),
		}
		err = c.UpdatePassword(u, &req)
		assertions.Nil(err)

		creds := models.User{
			Username: u.Username,
			Password: req.NewPassword,
		}
		_, err = c.Login(&creds)
		assertions.Nil(err)
	}))
	t.Run("Invalid Password", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)

		c, err := New(WithDB(db), WithSecret(random.Bytes(16)))
		assertions.Nil(err)

		u := models.RandomUser()
		err = c.Register(u)
		assertions.Nil(err)

		req := UpdatePasswordRequest{
			OldPassword: random.String(16),
			NewPassword: random.String(16),
		}
		err = c.UpdatePassword(u, &req)
		assertions.NotNil(err)
	}))
}
