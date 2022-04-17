package userdomain

import (
	users "swimming-content-management/data/user"
)

type UserRepository interface {
	SignUp(user *users.User) (*users.User, error)
	Login(user *users.User) (*users.AuthPayload, error)
}
