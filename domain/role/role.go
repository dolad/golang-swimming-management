package role

import (
	"swimming-content-management/domain/permission"
)

type Role struct {
	ID          uint32
	Name        string
	Permissions []*permission.Permission
}
