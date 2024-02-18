package database

import (
	"DebTour/models"
	"gorm.io/gorm"
)

func GetUserByUsername(username string, db *gorm.DB) (models.User, error) {
	var user models.User
	result := db.Model(&models.User{}).Where("username = ?", username).First(&user)
	return user, result.Error
}

func CreateUser(user *models.User, db *gorm.DB) error {
	result := db.Model(&models.User{}).Create(user)
	return result.Error
}

func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	result := db.Model(&models.User{}).Find(&users)
	return users, result.Error
}

func DeleteUser(user *models.User, db *gorm.DB) error {
	result := db.Model(&models.User{}).Where("username = ?", user.Username).Delete(user)
	return result.Error
}

func UpdateUser(user *models.User, db *gorm.DB) error {
	_, err := GetUserByUsername(user.Username, db)
	if err != nil {
		return err
	}
	result := db.Model(&models.User{}).Where("username = ?", user.Username).Updates(user)
	return result.Error
}
