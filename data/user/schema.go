package users

import (
	"log"
	"swimming-content-management/data/role"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Username    string    `gorm:"size:255;not null;index:unique" json:"username"`
	Email       string    `gorm:"index:unique; not null" json:"email" `
	Password    string    `json:"password" gorm:"not null"`
	Surname     string    `json:"surname" gorm:"not null"`
	FirstName   string    `json:"firstname" gorm:"not null"`
	DateofBirth time.Time `json:"dateofbirth" gorm:"not null"`
	PhoneNumber string    `json:"phonenumber" gorm:"not null"`
	Address     string    `json:"address" gorm:"not null"`
	PostCode    string    `json:"postcode gorm:not null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	RoleID      uint32
	Role        role.Role
}

func Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyHash(oldPassword string, newPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(oldPassword), []byte(newPassword))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("Id", uuid)
}
