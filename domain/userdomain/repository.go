package userdomain

import (
	uuid "github.com/satori/go.uuid"

	users "swimming-content-management/data/user"
)

type UserRepository interface {
	SignUp(user *users.User) (*users.User, error)
	Login(user *users.User) (*users.AuthPayload, error)
	GetUsers() ([]users.User, error)
	GetUserById(userId uuid.UUID) (*users.User, error)
	UpdateUserProfile(user *users.User, userId uuid.UUID) (*users.User, error)
}
