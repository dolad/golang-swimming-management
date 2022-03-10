package userdomain

type UserRepository interface {
	SignUp(user *User) (*User, error)
}
