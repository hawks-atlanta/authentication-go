package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/hawks-atlanta/authentication-go/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestBaseModel_BeforeSafe(t *testing.T) {
	t.Run("Succeed", database.Test(func(t *testing.T, db *gorm.DB) {
		db.AutoMigrate(&Model{})
		t.Run("Null UUID", func(t *testing.T) {
			assert.Nil(t, db.Create(&Model{}).Error)
		})
		t.Run("Set UUID", func(t *testing.T) {
			assert.Nil(t, db.Create(&Model{UUID: uuid.New()}).Error)
		})
	}))
}
