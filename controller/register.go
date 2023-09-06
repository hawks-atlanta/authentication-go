package controller

import (
	"github.com/hawks-atlanta/authentication-go/models"
)

func (c *Controller) Register(user *models.User) (err error) {
	user.Model = models.Model{}
	return c.DB.Create(user).Error
}
