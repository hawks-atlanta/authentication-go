package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Option func(r *Router)

func WithEngine(e *gin.Engine) Option {
	return func(r *Router) {
		r.Engine = e
	}
}

func WithDatabase(db *gorm.DB) Option {
	return func(r *Router) {
		r.DB = db
	}
}
