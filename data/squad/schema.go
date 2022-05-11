package squad

import (
	"github.com/jinzhu/gorm"
	users "swimming-content-management/data/user"
)

type Squad struct {
	gorm.Model
	Name     string `gorm:"size:255;not null" json:"name" json:"name,omitempty"`
	Coach    *users.User
	Swimmers []*users.User
}
