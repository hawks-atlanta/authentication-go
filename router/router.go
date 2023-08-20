package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	*gin.Engine
	DB *gorm.DB
}

func New(opts ...Option) *gin.Engine {
	var r Router
	for _, opt := range opts {
		opt(&r)
	}

	r.Any(EchoRoute, r.AnyEcho)

	return r.Engine
}
