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

	r.POST(LoginRoute, r.Login)
	r.POST(RegisterRoute, r.Register)
	// Authentication required
	_ = r.Group(RootRoute, r.Authorize)

	return r.Engine
}
