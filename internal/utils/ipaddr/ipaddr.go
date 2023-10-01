package ipaddr

import (
	"github.com/gin-gonic/gin"
)

func GetIpAddr(ctx *gin.Context) string {

	IPAddress := ctx.Request.Header.Get("X-Real-Ip") // Try to get the IP address of a client even if it's behind a proxy or a load balancer

	if IPAddress == "" {
		IPAddress = ctx.Request.Header.Get("X-Forwarded-For") // If first method get blank result, try to get the IP address through X-Forwarded-For header
	}
	if IPAddress == "" {
		IPAddress = ctx.Request.RemoteAddr // If the lasts methods don't get success, then get the IP from remote address request
	}

	return IPAddress
}
