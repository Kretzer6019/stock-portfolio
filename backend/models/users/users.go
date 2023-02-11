package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          int       `gorm:"id,omitempty"`
	Email       string    `gorm:"email,omitempty"`
	Password    string    `gorm:"password,omitempty"`
	DateCreated time.Time `gorm:"date_created,omitempty"`
}

func SelectUser(id int, db *gorm.DB) (User, error) {
	var user User
	err := db.First(&user, "id = ?", id).Error

	return user, err
}
