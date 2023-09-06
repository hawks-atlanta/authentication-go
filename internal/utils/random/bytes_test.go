package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytes(t *testing.T) {
	assert.NotEqual(t, Bytes(16), Bytes(16))
}
