package role

import (
	dataRole "swimming-content-management/data/role"
)

type RoleServices interface {
	CreateRole(role *dataRole.Role) (*dataRole.Role, error)
	FindByName(roleName string) (dataRole.Role, error)
	FindById(id uint32) (dataRole.Role, error)
	FindAll() ([]dataRole.Role, error)
}

// handles buisiness logic
type Service struct {
	repository RoleRepository
}

func (svc *Service) CreateRole(role *dataRole.Role) (*dataRole.Role, error) {
	return svc.repository.CreateRole(role)
}

func (svc *Service) FindByName(roleName string) (dataRole.Role, error) {
	return svc.repository.FindByName(roleName)
}

func (svc *Service) FindById(id uint32) (dataRole.Role, error) {
	return svc.repository.FindById(id)
}
func (svc *Service) FindAll() ([]dataRole.Role, error) {
	return svc.repository.FindAll()
}

func NewService(repository RoleRepository) *Service {
	return &Service{repository: repository}
}
