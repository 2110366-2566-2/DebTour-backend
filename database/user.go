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
	db.SavePoint("BeforeCreateUser")
	result := db.Model(&models.User{}).Create(user)
	if result.Error != nil {
		db.RollbackTo("BeforeCreateUser")
		return result.Error
	}
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
	result := db.Model(&models.User{}).Where("username = ?", username).Updates(user)
	return result.Error
}
