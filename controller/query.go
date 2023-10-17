package controller

import "github.com/hawks-atlanta/authentication-go/models"

type UserRequest struct {
	Username string `json:"username,omitempty"`
	UUID     string `json:"uuid,omitempty"`
}

func (c *Controller) UserByUsername(req *UserRequest) (user models.User, err error) {
	err = c.DB.
		Where("username = ?", req.Username).
		First(&user).
		Error
	return user, err
}

func (c *Controller) UsernameByUUID(req *UserRequest) (user models.User, err error) {
	err = c.DB.
		Where("uuid = ?", req.UUID).
		First(&user).
		Error
	return user, err
}
