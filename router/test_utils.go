package router

import (
	"net"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func NewDefault(t *testing.T, db *gorm.DB) (expect *httpexpect.Expect, closeFunc func()) {
	e := gin.New()
	e = New(WithEngine(e), WithDatabase(db))

	l, err := net.Listen("tcp", "127.0.0.1:0")
	assert.Nil(t, err)

	go e.RunListener(l)

	expect = httpexpect.Default(t, "http://"+l.Addr().String())
	closeFunc = func() {
		l.Close()
	}
	return expect, closeFunc
}
