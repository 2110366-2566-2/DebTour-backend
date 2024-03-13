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

func DeleteUserByUsername(username string, db *gorm.DB) error {
	result := db.Model(&models.User{}).Where("username = ?", username).Delete(&models.User{})
	return result.Error
}

func UpdateUserByUsername(username string, user models.User, db *gorm.DB) error {
	// Check if the user record exists
	existingUser, err := GetUserByUsername(username, db)
	if err != nil {
		return err
	}

	// Update the fields of the existing user record with the values from the provided user struct
	result := db.Model(&existingUser).Updates(user)
	return result.Error
}