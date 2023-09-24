package ipaddr

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetIpAddr(t *testing.T) {

	req := httptest.NewRequest("GET", "/echo", nil)
	req.Header.Set("X-Real-Ip", "192.168.1.1")
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ipAddress := GetIpAddr(ctx)
	assert.Equal(t, "192.168.1.1", ipAddress, "La dirección IP debe ser la misma que X-Real-Ip")

	req = httptest.NewRequest("GET", "/echo", nil)
	req.Header.Set("X-Forwarded-For", "192.168.2.2")
	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	ctx.Request = req
	ipAddress = GetIpAddr(ctx)
	assert.Equal(t, "192.168.2.2", ipAddress, "La dirección IP debe ser la misma que X-Forwarded-For")

	req = httptest.NewRequest("GET", "/echo", nil)
	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	ctx.Request = req
	ipAddress = GetIpAddr(ctx)
	assert.NotEmpty(t, ipAddress, "La dirección IP no debe estar vacía")
}
