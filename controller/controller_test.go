package controller

import (
	"fmt"
	"testing"

	"github.com/hawks-atlanta/authentication-go/database"
	"github.com/hawks-atlanta/authentication-go/internal/utils/random"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestNew(t *testing.T) {
	t.Run("Succeed", database.Test(func(t *testing.T, db *gorm.DB) {
		_, err := New(WithDB(db), WithSecret(random.Bytes(16)))
		assert.Nil(t, err)
	}))
	t.Run("Fail", database.Test(func(t *testing.T, db *gorm.DB) {
		_, err := New(func(c *Controller) error { return fmt.Errorf("intended failure") })
		assert.NotNil(t, err)
	}))
}
