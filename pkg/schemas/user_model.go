package schemas

import (
	"errors"

	"github.com/kasfil/offspace/pkg/utils"

	"gorm.io/gorm"
)

// User : user database struct
type User struct {
	gorm.Model
	ID       uint
	Name     string
	Username string  `gorm:"unique"`
	Email    *string `gorm:"unique"`
	Password string
}

// BeforeCreate : hook before a user is saved
func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	if u.Password == "" {
		err = errors.New("Password couldn't empty string")
	}

	// hash user password
	passwordHash, err := utils.GeneratePassword(u.Password)
	if err != nil {
		return
	}

	u.Password = passwordHash

	return
}
