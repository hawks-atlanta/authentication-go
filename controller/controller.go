package controller

import (
	"github.com/hawks-atlanta/authentication-go/models"
	"gorm.io/gorm"
)

type (
	Option     func(c *Controller) error
	Controller struct {
		*gorm.DB
	}
)

func WithDB(db *gorm.DB) Option {
	return func(c *Controller) error {
		c.DB = db
		return c.DB.AutoMigrate(&models.User{})
	}
}

func New(options ...Option) (c *Controller, err error) {
	c = new(Controller)
	for _, option := range options {
		err = option(c)
		if err != nil {
			break
		}
	}
	return c, err
}
