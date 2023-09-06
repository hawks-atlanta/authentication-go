package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert.NotEqual(t, String(16), String(16))
}
