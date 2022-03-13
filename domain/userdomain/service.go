package userdomain

type UserService interface {
	SignUp(user *User) (*User, error)
	Login(user *User) (*AuthPayload, error)
}

// handles buisiness logic
type Service struct {
	repository UserRepository
}

func (svc *Service) SignUp(user *User) (*User, error) {
	return svc.repository.SignUp(user)
}

func (svc *Service) Login(user *User) (*AuthPayload, error) {
	return svc.repository.Login(user)
}

func NewService(repository UserRepository) *Service {
	return &Service{repository: repository}
}
