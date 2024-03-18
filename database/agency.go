package database

import (
	"DebTour/models"

	"gorm.io/gorm"
)

func GetAgencyByUsername(username string, db *gorm.DB) (models.Agency, error) {
	var agency models.Agency
	result := db.Model(&models.Agency{}).Where("username = ?", username).First(&agency)

	return agency, result.Error
}

func CreateAgency(agency *models.Agency, db *gorm.DB) error {
	result := db.Model(&models.Agency{}).Create(agency)
	return result.Error
}

func GetAllAgencies(db *gorm.DB) ([]models.Agency, error) {
	var agencies []models.Agency
	result := db.Model(&models.Agency{}).Find(&agencies)
	return agencies, result.Error
}

func DeleteAgencyByUsername(username string, db *gorm.DB) error {
	result := db.Model(&models.Agency{}).Where("username = ?", username).Delete(&models.Agency{})
	return result.Error
}

func UpdateAgencyByUsername(username string, agency models.Agency, db *gorm.DB) error {
	existingUser, err := GetAgencyByUsername(username, db)
	if err != nil {
		return err
	}
	result := db.Model(&existingUser).Where("username = ?", username).Updates(agency)
	return result.Error
}
