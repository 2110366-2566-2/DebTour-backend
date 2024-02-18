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
