package permission

import "github.com/jinzhu/gorm"

type Permission struct {
	gorm.Model
	Name string `gorm:"size:255;not null" json:"name"`
}
