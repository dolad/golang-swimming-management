package role

import (
	dataRole "swimming-content-management/data/role"
)

type RoleRepository interface {
	CreateRole(role *dataRole.Role) (*dataRole.Role, error)
	FindByName(roleName string) (dataRole.Role, error)
	FindById(id uint32) (dataRole.Role, error)
	FindAll() ([]dataRole.Role, error)
}
