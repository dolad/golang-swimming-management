package role

import (
	"github.com/jinzhu/gorm"
	"swimming-content-management/data/permission"
)

type Role struct {
	gorm.Model
	Name        string                   `gorm:"size:255;not null" json:"name" json:"name,omitempty"`
	Permissions []*permission.Permission `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
}
