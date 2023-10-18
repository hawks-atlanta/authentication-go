package controller

import (
	"fmt"

	"github.com/hawks-atlanta/authentication-go/models"
)

var (
	ErrDuplicatedUSer = fmt.Errorf("constraint failed")
)

func (c *Controller) Register(user *models.User) (err error) {
	user.Model = models.Model{}
	return c.DB.Create(user).Error
}
