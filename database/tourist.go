package database

import (
	"DebTour/models"
	"gorm.io/gorm"
)

func GetTouristByUsername(username string, db *gorm.DB) (models.Tourist, error) {
	var tourist models.Tourist
	result := db.Model(&models.Tourist{}).First(&tourist, username)

	return tourist, result.Error
}


func CreateTourist(tourist *models.Tourist, db *gorm.DB) error {
    result := db.Model(&models.Tourist{}).Create(tourist)
    return result.Error
}

func GetAllTourists(db *gorm.DB) ([]models.Tourist, error) {
	var tourists []models.Tourist
	result := db.Model(&models.Tourist{}).Find(&tourists)
	return tourists, result.Error
}

func DeleteTourist(tourist models.Tourist, db *gorm.DB) error {
	result := db.Model(&models.Tourist{}).Where("username = ?", tourist.Username).Delete(tourist)
	return result.Error
}

func UpdateTourist(tourist models.Tourist, db *gorm.DB) error {
	_, err := GetTouristByUsername(tourist.Username, db)
	if err != nil {
		return err
	}
	result := db.Model(&models.Tourist{}).Where("username = ?", tourist.Username).Updates(tourist)
	return result.Error
}
