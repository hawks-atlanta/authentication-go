package controller

import (
	"testing"

	"github.com/hawks-atlanta/authentication-go/database"
	"github.com/hawks-atlanta/authentication-go/internal/utils/jwt"
	"github.com/hawks-atlanta/authentication-go/internal/utils/random"
	"github.com/hawks-atlanta/authentication-go/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestController_Authorize(t *testing.T) {
	t.Run("Succeed", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)

		c, err := New(WithDB(db), WithSecret(random.Bytes(16)))
		assertions.Nil(err)

		u := models.RandomUser()
		err = c.Register(u)
		assertions.Nil(err)

		token := c.JWT.New(u.Claims())

		user, err := c.Authorize(token)
		assertions.Nil(err)

		assertions.Equal(u.UUID, user.UUID)
		assertions.Equal(u.Username, user.Username)
	}))
	t.Run("Invalid Token", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)

		c, err := New(WithDB(db), WithSecret(random.Bytes(16)))
		assertions.Nil(err)

		_, err = c.Authorize("INVALID")
		assertions.NotNil(err)
	}))
	t.Run("Invalid SECRET", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)

		c, err := New(WithDB(db), WithSecret(random.Bytes(16)))
		assertions.Nil(err)

		_, err = c.Authorize(jwt.New(random.Bytes(16)).New(models.RandomUser().Claims()))
		assertions.NotNil(err)
	}))
	t.Run("Invalid User", database.Test(func(t *testing.T, db *gorm.DB) {
		assertions := assert.New(t)

		c, err := New(WithDB(db), WithSecret(random.Bytes(16)))
		assertions.Nil(err)

		_, err = c.Authorize(c.JWT.New(models.RandomUser().Claims()))
		assertions.NotNil(err)
	}))
}
