package userdomain

type UserService interface {
	SignUp(user *User) (*User, error)
}

// handles buisiness logic
type Service struct {
	repository UserRepository
}

func (svc *Service) SignUp(user *User) (*User, error) {
	return svc.repository.SignUp(user)
}

func NewService(repository UserRepository) *Service {
	return &Service{repository: repository}
}
