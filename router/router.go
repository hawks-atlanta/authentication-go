package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hawks-atlanta/authentication-go/controller"
)

type Router struct {
	*gin.Engine
	C *controller.Controller
}

func New(opts ...Option) *gin.Engine {
	var r Router
	for _, opt := range opts {
		opt(&r)
	}

	r.Any(EchoRoute, r.AnyEcho)
	r.POST(LoginRoute, r.Login)

	return r.Engine
}
