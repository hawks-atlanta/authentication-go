package models

import (
	"testing"

	"github.com/hawks-atlanta/authentication-go/database"
	"github.com/hawks-atlanta/authentication-go/internal/utils/random"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestRandomUser(t *testing.T) {
	assert.NotEqual(t, RandomUser(), RandomUser())
}

func TestUser(t *testing.T) {
	t.Run("Succeed", database.Test(func(t *testing.T, db *gorm.DB) {
		assert.Nil(t, db.AutoMigrate(&User{}))

		t.Run("ValidUser", func(t *testing.T) {
			assertions := assert.New(t)
			u := RandomUser()
			assertions.Nil(db.Create(u).Error)
		})
		t.Run("RepeatedUser", func(t *testing.T) {
			assertions := assert.New(t)
			u := RandomUser()
			assertions.Nil(db.Create(u).Error)
			assertions.NotNil(db.Create(u).Error)
		})
		t.Run("Claims", func(t *testing.T) {
			assertions := assert.New(t)
			u := RandomUser()
			assertions.Nil(db.Create(u).Error)
			assertions.NotNil(u.Claims())
		})
		t.Run("FromClaims-Valid", func(t *testing.T) {
			assertions := assert.New(t)
			u := RandomUser()
			var u2 User
			assertions.Nil(u2.FromClaims(u.Claims()))
			assertions.Equal(u.UUID, u2.UUID)
		})
		t.Run("FromClaims-NoUUID", func(t *testing.T) {
			assertions := assert.New(t)
			u := RandomUser()
			var u2 User
			claims := u.Claims()
			delete(claims, "uuid")
			assertions.NotNil(u2.FromClaims(claims))
		})
		t.Run("Authenticate-Succeed", func(t *testing.T) {
			assertions := assert.New(t)
			u := RandomUser()
			u.PasswordHash, _ = bcrypt.GenerateFromPassword([]byte(u.Password), DefaultPasswordCost)
			assertions.True(u.Authenticate(u.Password))
		})
		t.Run("Authenticate-Fail", func(t *testing.T) {
			assertions := assert.New(t)
			u := RandomUser()
			u.PasswordHash, _ = bcrypt.GenerateFromPassword([]byte(u.Password), DefaultPasswordCost)
			assertions.False(u.Authenticate("wrong"))
		})
		t.Run("BeforeSave-ValidUser", func(t *testing.T) {
			assertions := assert.New(t)
			u := RandomUser()
			assertions.Nil(db.Create(u).Error)
		})
		t.Run("BeforeSave-NoPassword", func(t *testing.T) {
			assertions := assert.New(t)
			u := RandomUser()
			u.Password = ""
			assertions.NotNil(db.Create(u).Error)
		})
		t.Run("Update - Password", func(t *testing.T) {
			assertions := assert.New(t)
			u := RandomUser()
			assertions.Nil(db.Create(u).Error)
			update := User{
				Password: random.String(16),
			}
			assertions.Nil(db.
				Where("uuid = ?", u.UUID).
				Updates(&update).
				Error)
			var user User
			assertions.Nil(db.
				Where("uuid = ? AND password_hash = ?", u.UUID, update.PasswordHash).
				First(&user).
				Error)
			assertions.Equal(u.UUID, user.UUID)
			assertions.Equal(u.Username, user.Username)
			assertions.Equal(update.PasswordHash, user.PasswordHash)
		})
	}))
}
