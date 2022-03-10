package users

import (
	userDomain "swimming-content-management/domain/userdomain"

	"github.com/gin-gonic/gin"
)

type UserRequestValidator struct {
	Email    string `binding:"required" json:"email"`
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}

func Bind(c *gin.Context) (*userDomain.User, error) {
	var json UserRequestValidator
	if err := c.ShouldBindJSON(&json); err != nil {
		return nil, err
	}

	newUser := &userDomain.User{
		Email:    json.Email,
		Username: json.Username,
		Password: json.Password,
	}

	return newUser, nil
}
