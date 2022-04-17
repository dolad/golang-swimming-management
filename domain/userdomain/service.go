package userdomain

import users "swimming-content-management/data/user"

type UserService interface {
	SignUp(user *users.User) (*users.User, error)
	Login(user *users.User) (*users.AuthPayload, error)
}

// handles buisiness logic
type Service struct {
	repository UserRepository
}

func (svc *Service) SignUp(user *users.User) (*users.User, error) {
	return svc.repository.SignUp(user)
}

func (svc *Service) Login(user *users.User) (*users.AuthPayload, error) {
	return svc.repository.Login(user)
}

func NewService(repository UserRepository) *Service {
	return &Service{repository: repository}
}
