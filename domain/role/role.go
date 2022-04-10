package role

import (
	"swimming-content-management/domain/permission"
)

type Role struct {
	ID          uint
	Name        string
	Permissions []*permission.Permission
}
