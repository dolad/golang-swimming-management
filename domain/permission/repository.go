package permission

import dataPermission "swimming-content-management/data/permission"

type PermissionRepository interface {
	CreatePermission(permission *dataPermission.Permission) (*dataPermission.Permission, error)
	FindByName(permissionName string) (*dataPermission.Permission, error)
	FindById(id uint32) (*dataPermission.Permission, error)
	FindAll() ([]dataPermission.Permission, error)
}
