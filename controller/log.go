package controller

import (
	"github.com/hawks-atlanta/authentication-go/models"
)

func (c *Controller) Log(log *models.Log) (err error) {
	log.Model = models.Model{}
	return c.DB.Create(log).Error
}

func (c *Controller) GetLogs() (logs []models.Log, err error) {

	err = c.DB.Find(&logs).Error
	return logs, err
}
