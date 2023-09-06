package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hawks-atlanta/authentication-go/controller"
)

type Option func(r *Router)

func WithEngine(e *gin.Engine) Option {
	return func(r *Router) {
		r.Engine = e
	}
}

func WithController(c *controller.Controller) Option {
	return func(r *Router) {
		r.C = c
	}
}
