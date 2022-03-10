package users

import (
	"time"
)

// "github.com/jinzhu/gorm"
// "github.com/pkg/errors"

type User struct {
	Id        int       `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"unique; not null" json:"email" `
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// func Hash(password string) ([]byte, error) {
// 	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// }

// func VerifyPassword(hashedPassword, password string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// }

// // before create hooks it runs before create functions
// func (u *User) BeforeSave() error {
// 	hashedPassword, err := Hash(u.Password)
// 	if err != nil {
// 		return err
// 	}
// 	u.Password = string(hashedPassword)
// 	return nil
// }

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
