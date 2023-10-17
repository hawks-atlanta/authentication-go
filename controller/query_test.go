package controller

import (
	"testing"

	"github.com/hawks-atlanta/authentication-go/database"
	"github.com/hawks-atlanta/authentication-go/internal/utils/random"
	"github.com/hawks-atlanta/authentication-go/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUserByUsername(t *testing.T) {
	t.Run("Succeed", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)

		c, err := New(WithDB(db), WithSecret(random.String(16)))
		assertions.Nil(err)

		u := models.RandomUser()
		err = c.Register(u)
		assertions.Nil(err)

		req := UserRequest{
			Username: *u.Username,
		}
		ou, err := c.UserByUsername(&req)

		assertions.Equal(u.UUID, ou.UUID)
	}))
}

func TestUsernameByUUID(t *testing.T) {
	t.Run("Succeed", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)

		c, err := New(WithDB(db), WithSecret(random.String(16)))
		assertions.Nil(err)

		u := models.RandomUser()
		err = c.Register(u)
		assertions.Nil(err)

		req := UserRequest{
			UUID: u.UUID.String(),
		}
		ou, err := c.UsernameByUUID(&req)

		assertions.Equal(u.Username, ou.Username)
	}))
}
