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

func CreateUser(user *User) error {
	result := db.Model(&User{}).Create(user)
	return result.Error
}

func GetAllUsers() ([]User, error) {
	var users []User
	result := db.Model(&User{}).Find(&users)
	return users, result.Error
}

func DeleteUser(user *User) error {
	result := db.Model(&User{}).Where("username = ?", user.Username).Delete(user)
	return result.Error
}

func UpdateUser(user *User) error {
	_, err := GetUserByUsername(user.Username)
	if err != nil {
		return err
	}
	result := db.Model(&User{}).Where("username = ?", user.Username).Updates(user)
	return result.Error	
}
