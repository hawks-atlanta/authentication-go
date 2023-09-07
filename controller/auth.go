package controller

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hawks-atlanta/authentication-go/models"
	"gorm.io/gorm"
)

func (c *Controller) Authorize(jwtS string) (user models.User, err error) {
	tok, err := jwt.Parse(jwtS, c.JWT.KeyFunc)
	if err != nil {
		err = ErrUnauthorized
		return user, err
	}
	var check models.User
	err = check.FromClaims(tok.Claims.(jwt.MapClaims))
	if err != nil {
		err = ErrUnauthorized
		return user, err
	}
	err = c.DB.
		Where("uuid = ?", check.UUID).
		First(&user).
		Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = ErrUnauthorized
	}
	return user, err
}
