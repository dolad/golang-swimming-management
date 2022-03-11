package userdomain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	Id          uuid.UUID
	Username    string
	Email       string
	Password    string
	Surname     string
	FirstName   string
	DateofBirth time.Time
	PhoneNumber string
	Address     string
	PostCode    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
