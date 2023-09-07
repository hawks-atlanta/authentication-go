package controller

import (
	"github.com/hawks-atlanta/authentication-go/models"
)

type UpdatePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func (c *Controller) UpdatePassword(session *models.User, req *UpdatePasswordRequest) (err error) {
	if !session.Authenticate(req.OldPassword) {
		return ErrUnauthorized
	}
	update := models.User{
		Model:    session.Model,
		Password: req.NewPassword,
	}
	err = c.DB.
		Where("uuid = ? AND password_hash = ?", session.UUID, session.PasswordHash).
		Updates(&update).
		Error
	return err
}
