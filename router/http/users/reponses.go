package users

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserResponse struct {
	Id          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Surname     string    `json:"surname"`
	FirstName   string    `json:"firstname"`
	DateofBirth time.Time `json:"dateofbirth"`
	PhoneNumber string    `json:"phonenumber"`
	Address     string    `json:"address"`
	PostCode    string    `json:"postcode"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ListUserResponse struct {
	Data []UserResponse `json:"data"`
}
type AuthPayloadResponse struct {
	Id          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Surname     string    `json:"surname"`
	FirstName   string    `json:"firstname"`
	DateofBirth time.Time `json:"dateofbirth"`
	PhoneNumber string    `json:"phonenumber"`
	Address     string    `json:"address"`
	PostCode    string    `json:"postcode"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Token       string    `json:"accessToken"`
}
