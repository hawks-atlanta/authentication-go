package jwt

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hawks-atlanta/authentication-go/models"
	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		j := New([]byte("TESTING"))
		jwtTok := j.New(jwt.MapClaims{"uuid": "my_uuid"})
		assert.Greater(t, len(jwtTok), 0)
		tok, err := jwt.Parse(jwtTok, j.KeyFunc)
		assert.Nil(t, err)
		assert.True(t, tok.Valid)
	})
	t.Run("User", func(t *testing.T) {
		j := New([]byte("TESTING"))
		user := models.RandomUser()
		user.UUID = uuid.New()
		jwtTok := j.New(user.Claims())
		assert.Greater(t, len(jwtTok), 0)
		tok, err := jwt.Parse(jwtTok, j.KeyFunc)
		assert.Nil(t, err)
		assert.True(t, tok.Valid)
		var user2 models.User
		assert.Nil(t, user2.FromClaims(tok.Claims.(jwt.MapClaims)))
		assert.Equal(t, user.UUID, user2.UUID)
	})
}
