package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strings"
	users "swimming-content-management/data/user"
	"time"
)

type fieldError struct {
	err validator.FieldError
}

type UserRequestValidator struct {
	Email       string    `json:"email" binding:"required" validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	Username    string    `json:"username" binding:"required" validate:"min=1,max=16,regexp=^[a-zA-Z]*$"`
	Password    string    `json:"password" binding:"required" validate:"min=6,max=10"`
	Surname     string    `json:"surname" binding:"required" validate:"min=1,max=20"`
	FirstName   string    `json:"firstname" binding:"required" validate:"min=1,max=20"`
	DateofBirth time.Time `json:"dateofbirth"`
	PhoneNumber string    `json:"phonenumber" binding:"required" validate:"min=10,max=11"`
	Address     string    `json:"address" binding:"required" validate:"min=5,max=60"`
	PostCode    string    `json:"postcode" binding:"required" validate:"min=4,max=10"`
	RoleId      uint32    `json:"role_id" binding:"required" `
}

type UserAuthRequestValidator struct {
	Email    string `json:"email" binding:"required" validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	Password string `json:"password" binding:"required" validate:"min=6,max=10"`
}

func LoginRequestValidator(c *gin.Context) (*users.User, string) {
	var json UserAuthRequestValidator

	if err := c.ShouldBind(&json); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			return nil, fieldError{fieldErr}.String() // exit on first error
		}
	}

	authUser := &users.User{
		Email:    json.Email,
		Password: json.Password,
	}

	return authUser, ""

}

func Bind(c *gin.Context) (*users.User, error) {
	var json UserRequestValidator
	if err := c.ShouldBindJSON(&json); err != nil {
		return nil, err
	}

	newUser := &users.User{
		Email:       json.Email,
		Username:    json.Username,
		Password:    json.Password,
		Surname:     json.Surname,
		FirstName:   json.FirstName,
		DateofBirth: json.DateofBirth,
		PhoneNumber: json.PhoneNumber,
		Address:     json.Address,
		PostCode:    json.PostCode,
		RoleID:      json.RoleId,
	}

	return newUser, nil
}

func (q fieldError) String() string {
	var sb strings.Builder

	sb.WriteString("validation failed on field '" + q.err.Field() + "'")
	sb.WriteString(", condition: " + q.err.ActualTag())

	// Print condition parameters, e.g. oneof=red blue -> { red blue }
	if q.err.Param() != "" {
		sb.WriteString(" { " + q.err.Param() + " }")
	}

	if q.err.Value() != nil && q.err.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", q.err.Value()))
	}

	return sb.String()
}
