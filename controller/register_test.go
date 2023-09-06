package controller

import (
	"testing"

	"github.com/hawks-atlanta/authentication-go/database"
	"github.com/hawks-atlanta/authentication-go/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestController_Register(t *testing.T) {
	t.Run("Succeed", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)

		c, err := New(WithDB(db))
		assertions.Nil(err)

		u := models.RandomUser()
		assertions.Nil(c.Register(u))

		user, err := c.Login(u)
		assertions.Nil(err)

		assertions.Equal(u.UUID, user.UUID)
	}))
	t.Run("Repeated username", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)

		c, err := New(WithDB(db))
		assertions.Nil(err)

		u := models.RandomUser()
		assertions.Nil(c.Register(u))
		assertions.NotNil(c.Register(u))
	}))
}
