package router

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InternalServerError(t *testing.T) {
	err := fmt.Errorf("ERROR")
	r := InternalServerError(err)
	assert.Equal(t, err.Error(), r.Message)
}
