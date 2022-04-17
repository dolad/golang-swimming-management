package permission

import dataPermission "swimming-content-management/data/permission"

type PermissionService interface {
	CreatePermission(permission *dataPermission.Permission) (*dataPermission.Permission, error)
	FindByName(permissionName string) (*dataPermission.Permission, error)
	FindById(id uint32) (*dataPermission.Permission, error)
	FindAll() ([]dataPermission.Permission, error)
}

// handles buisiness logic
type Service struct {
	repository PermissionRepository
}

func (svc *Service) CreatePermission(permission *dataPermission.Permission) (*dataPermission.Permission, error) {
	return svc.repository.CreatePermission(permission)
}

func (svc *Service) FindByName(permissionName string) (*dataPermission.Permission, error) {
	return svc.repository.FindByName(permissionName)
}

func (svc *Service) FindById(id uint32) (*dataPermission.Permission, error) {
	return svc.repository.FindById(id)
}
func (svc *Service) FindAll() ([]dataPermission.Permission, error) {
	return svc.repository.FindAll()
}

func NewService(repository PermissionRepository) *Service {
	return &Service{repository: repository}
}
