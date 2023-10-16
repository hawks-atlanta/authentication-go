package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomLog(t *testing.T) {
	firstLog, _ := RandomLog()
	secondLog, _ := RandomLog()
	assert.NotEqual(t, firstLog, secondLog)
}
