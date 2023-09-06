package controller

import (
	"errors"
	"fmt"

	"github.com/hawks-atlanta/authentication-go/models"
	"gorm.io/gorm"
)

var (
	ErrUnauthorized = fmt.Errorf("unauthorized")
)

func (c *Controller) Login(creds *models.User) (user models.User, err error) {
	if creds.Username == nil {
		err = ErrUnauthorized
		return user, err
	}
	err = c.DB.
		Where("username = ?", *creds.Username).
		First(&user).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("auth failed: %w -> %w", err, ErrUnauthorized)
			return user, err
		}
		err = fmt.Errorf("something went wrong: %w", err)
		return user, err
	}
	if !user.Authenticate(creds.Password) {
		err = ErrUnauthorized
	}
	return user, err
}
