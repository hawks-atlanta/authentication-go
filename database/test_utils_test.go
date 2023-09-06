package database

import (
	"testing"

	"gorm.io/gorm"
)

func Test_Test(t *testing.T) {
	t.Run("Succeed", Test(func(t *testing.T, d *gorm.DB) {

	}))
}
