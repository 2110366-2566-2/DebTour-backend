package database

import (
	"DebTour/models"

	"gorm.io/gorm"
)

func GetTouristByUsername(username string, db *gorm.DB) (tourist models.Tourist, err error) {
	result := db.Model(&models.Tourist{}).Where("username = ?", username).First(&tourist)

	return tourist, result.Error
}

func CreateTourist(tourist *models.Tourist, db *gorm.DB) (err error) {
	result := db.Model(&models.Tourist{}).Create(tourist)
	return result.Error
}

func GetAllTourists(db *gorm.DB) (tourists []models.Tourist, err error) {
	result := db.Model(&models.Tourist{}).Find(&tourists)
	return tourists, result.Error
}

// func delete tourist by username and db
func DeleteTouristByUsername(username string, db *gorm.DB) (err error) {
	result := db.Model(&models.Tourist{}).Where("username = ?", username).Delete(&models.Tourist{})
	return result.Error
}

func UpdateTouristByUsername(username string, tourist models.Tourist, db *gorm.DB) (err error){
	existingUser, err := GetTouristByUsername(username, db)
	if err != nil {
		return err
	}
	result := db.Model(&existingUser).Where("username = ?", username).Updates(tourist)
	return result.Error
}
