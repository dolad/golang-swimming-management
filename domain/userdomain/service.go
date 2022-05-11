package userdomain

import (
	uuid "github.com/satori/go.uuid"
	users "swimming-content-management/data/user"
)

type UserService interface {
	SignUp(user *users.User) (*users.User, error)
	Login(user *users.User) (*users.AuthPayload, error)
	GetUsers() ([]users.User, error)
	GetUserById(userId uuid.UUID) (*users.User, error)
	UpdateUserProfile(user *users.User, userId uuid.UUID) (*users.User, error)
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

func (svc *Service) GetUsers() ([]users.User, error) {
	return svc.repository.GetUsers()
}
func (svc *Service) GetUserById(userId uuid.UUID) (*users.User, error) {
	return svc.repository.GetUserById(userId)
}

func (svc *Service) UpdateUserProfile(user *users.User, userId uuid.UUID) (*users.User, error) {
	return svc.repository.UpdateUserProfile(user, userId)
}

func NewService(repository UserRepository) *Service {
	return &Service{repository: repository}
}
