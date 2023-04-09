package users

import (
	"dependencies/clients/db"
	"time"
)

type User struct {
	ID          int       `gorm:"id,omitempty"`
	Email       string    `gorm:"email,omitempty"`
	Password    string    `gorm:"password,omitempty"`
	DateCreated time.Time `gorm:"date_created,omitempty"`
}

func SelectUser(id int) (User, error) {
	var user User
	err := db.DB.First(&user, "id = ?", id).Error

	return user, err
}

func SelectUserByEmail(email string) (User, error) {
	var user User
	err := db.DB.First(&user, "email = ?", email).Error

	return user, err
}

func InsertUser(email string, password string) (int, error) {
	user := User{
		Email:    email,
		Password: password,
	}

	result := db.DB.Select("email", "password").Create(&user)

	return user.ID, result.Error
}
