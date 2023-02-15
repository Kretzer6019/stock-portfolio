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

func SelectUserByEmail(email string, db *gorm.DB) (User, error) {
	var user User
	err := db.First(&user, "email = ?", email).Error

	return user, err
}

func InsertUser(email string, password string, db *gorm.DB) (int, error) {
	user := User{
		Email:    email,
		Password: password,
	}

	result := db.Select("email", "password").Create(&user)

	return user.ID, result.Error
}
