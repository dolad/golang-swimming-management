package users

import (
	userDomain "swimming-content-management/domain/userdomain"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

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
}

type UserAuthRequestValidator struct {
	Email    string `json:"email" binding:"required" validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	Password string `json:"password" binding:"required" validate:"min=6,max=10"`
}

func LoginRequestValidator(c *gin.Context) (*userDomain.User, error) {
	var json UserAuthRequestValidator
	if err := c.ShouldBindJSON(&json); err != nil {
		return nil, err
	}

	if err := validator.Validate(&json); err != nil {
		return nil, err
	}

	authUser := &userDomain.User{
		Email:    json.Email,
		Password: json.Password,
	}

	return authUser, nil

}

func Bind(c *gin.Context) (*userDomain.User, error) {
	var json UserRequestValidator
	if err := c.ShouldBindJSON(&json); err != nil {
		return nil, err
	}

	if err := validator.Validate(&json); err != nil {
		return nil, err
	}

	newUser := &userDomain.User{
		Email:       json.Email,
		Username:    json.Username,
		Password:    json.Password,
		Surname:     json.Surname,
		FirstName:   json.FirstName,
		DateofBirth: json.DateofBirth,
		PhoneNumber: json.PhoneNumber,
		Address:     json.Address,
		PostCode:    json.PostCode,
	}

	return newUser, nil
}
