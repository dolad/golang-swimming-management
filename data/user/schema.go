package users

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// "github.com/jinzhu/gorm"
// "github.com/pkg/errors"

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
}

func Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// func VerifyPassword(hashedPassword, password string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// }

// // before create hooks it runs before create functions
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("Id", uuid)
}

// func (u *User) Prepare() {
// 	u.ID = 0
// 	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
// 	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
// 	u.CreatedAt = time.Now()
// 	u.UpdatedAt = time.Now()
// }

// func (u *User) SaveUser(db *gorm.DB) (*User, error) {
// 	var err error
// 	err = db.Debug().Create(&u).Error
// 	if err != nil {
// 		return &User{}, err
// 	}
// 	return u, nil
// }
