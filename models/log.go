package models

import (
	"math/rand"
	"time"

	"github.com/ddosify/go-faker/faker"
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

func RandomLog() (*Log, string) {
	user := RandomUser()
	actions := []string{"User login", "User registration", "User JWT renewal", "User password update", "Got user by username"}
	return &Log{
		User:      user,
		UserUUID:  user.UUID,
		Action:    actions[rand.Intn(4)],
		IpAddress: faker.NewFaker().RandomIP(),
	}, *user.Username
}
