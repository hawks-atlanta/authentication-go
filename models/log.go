package models

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	Model
	User      *User     `json:"user,omitempty" gorm:"foreignKey:UserUUID;constraint:OnUpdate:CASCADE"`
	UserUUID  uuid.UUID `json:"userUUID" gorm:"not null;"`
	Action    string    `json:"action" gorm:"not null"`
	IpAddress string    `json:"ipaddr" gorm:"not null"`
	LogTime   time.Time `json:"logTime" gorm:"not null;default:CURRENT_TIMESTAMP"`
}
