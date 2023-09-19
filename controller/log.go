package controller

import (
	"fmt"
	"time"

	"github.com/hawks-atlanta/authentication-go/models"
	"gorm.io/gorm"
)

type Filter[T any] struct {
	Object       T   `json:"object"`
	ItemsPerPage int `json:"itemsPerPage"`
	Page         int `json:"page"`
}

func (c *Controller) CreateLog(log *models.Log) (err error) {
	log.Model = models.Model{}
	return c.DB.Create(log).Error
}

func (c *Controller) GetLogs() (logs []models.Log, err error) {

	err = c.DB.Find(&logs).Error
	return logs, err
}

func (c *Controller) GetLogsByUser(filter *Filter[models.User]) (logs []models.Log, err error) {
	err = c.DB.Transaction(func(tx *gorm.DB) error {
		var user models.User
		err = tx.
			Where("username = ?", filter.Object.Username).
			First(&user).
			Error
		if err != nil {
			err = fmt.Errorf("failed to query user: %w", err)
			return err
		}

		err = tx.
			Limit(filter.ItemsPerPage).
			Offset((filter.Page-1)*filter.ItemsPerPage).
			Where("user_uuid = ?", user.UUID).
			Find(&logs).
			Error
		if err != nil {
			err = fmt.Errorf("failed to query user logs: %w", err)
		}
		return err
	})
	return logs, err
}

func (c *Controller) GetLogsByDate(filter *Filter[time.Time]) (logs []models.Log, err error) {

	err = c.DB.
		Limit(filter.ItemsPerPage).
		Offset((filter.Page-1)*filter.ItemsPerPage).
		Where("log_time > ?", filter.Object).
		Find(&logs).
		Error

	return logs, err
}
