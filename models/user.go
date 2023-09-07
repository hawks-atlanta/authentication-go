package models

import (
	"fmt"
	"regexp"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hawks-atlanta/authentication-go/internal/utils/random"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Model
	Username     *string `json:"username,omitempty" gorm:"unique;not null;"`
	Password     string  `json:"password,omitempty" gorm:"-"`
	PasswordHash []byte  `json:"-" gorm:"not null;"`
}

const (
	DefaultPasswordCost    = 12
	DefaultExpirationDelta = 30 * 24 * time.Hour
)

func (u *User) Claims() jwt.MapClaims {
	return jwt.MapClaims{
		"uuid": u.UUID.String(),
		"exp":  time.Now().Add(DefaultExpirationDelta).Unix(),
	}
}

func (u *User) FromClaims(m jwt.MapClaims) error {
	userUUID, ok := m["uuid"]
	if !ok {
		return fmt.Errorf("incomplete UUID")
	}
	u.UUID, _ = uuid.Parse(userUUID.(string))
	return nil
}

func (u *User) Authenticate(password string) bool {
	return u.PasswordHash != nil && bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password)) == nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if len(u.Password) != 0 {
		u.PasswordHash, err = bcrypt.GenerateFromPassword([]byte(u.Password), DefaultPasswordCost)
	}
	return err
}

var phoneRegex = regexp.MustCompile(`(?m)^\d{2,18}$`)

var (
	True  = true
	False = false
)

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	err = u.Model.BeforeCreate(tx)
	if err == nil && len(u.PasswordHash) == 0 {
		err = fmt.Errorf("password is empty")
	}
	return err
}

func RandomUser() *User {
	person := gofakeit.NewCrypto().Person()
	username := fmt.Sprint(person.FirstName, person.LastName, person.Gender, person.Hobby, person.Job.Company)
	return &User{
		Username: &username,
		Password: random.String(16),
	}
}
