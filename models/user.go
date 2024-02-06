package models

import (
	"time"
)

type User struct {
	Username    string    `gorm:"primary_key" json:"username"`
	Password    string    `json:"password"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	Image       string    `json:"image"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
}

func GetUserByUsername(username string) (User, error) {
	var user User
	result := db.Model(&User{}).Where("username = ?", username).First(&user)
	return user, result.Error
}

func CreateUser(username, password, phone, email, image string) error {
	user := User{
		Username: username,
		Password: password,
		Phone:    phone,
		Email:    email,
		Image:    image,
	}
	result := db.Model(&User{}).Create(&user)
	return result.Error
}

func GetAllUsers() ([]User, error) {
	var users []User
	result := db.Model(&User{}).Find(&users)
	return users, result.Error
}

func DeleteUser(username string) error {
	result := db.Model(&User{}).Where("username = ?", username).Delete(&User{})
	return result.Error
}

func UpdateUser(username, password, phone, email, image string) error {
	result := db.Model(&User{}).Where("username = ?", username).Updates(User{Password: password, Phone: phone, Email: email, Image: image})
	return result.Error
}
