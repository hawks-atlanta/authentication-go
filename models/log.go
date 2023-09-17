package models

import (
	"github.com/google/uuid"
)

type Log struct {
	Model
	User      User      `json:"user" gorm:"foreignKey:UserUUID;constraint:OnUpdate:CASCADE"`
	UserUUID  uuid.UUID `json:"userUUID" gorm:"uniqueIndex:idx_unique_station;not null;"`
	Action    string    `json:"action" gorm:"not null"`
	IpAddress string    `json:"ipaddr" gorm:"not null"`
}
