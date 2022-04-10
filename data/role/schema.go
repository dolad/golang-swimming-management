package role

import (
	"swimming-content-management/data/permission"
)

type Role struct {
	ID          uint                     `gorm:"primaryKey" json:"id"`
	Name        string                   `gorm:"size:255;not null" json:"name"`
	Permissions []*permission.Permission `gorm:"many2many:role_permissions"`
}
