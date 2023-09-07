package controller

import "github.com/hawks-atlanta/authentication-go/models"

type UserRequest struct {
	Username string `json:"username"`
}

func (c *Controller) UserByUsername(req *UserRequest) (user models.User, err error) {
	err = c.DB.
		Where("username = ?", req.Username).
		First(&user).
		Error
	return user, err
}
