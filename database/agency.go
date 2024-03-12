package database

import (
	"DebTour/models"

	"gorm.io/gorm"
)

func GetAgencyByUsername(username string, db *gorm.DB) (models.Agency, error) {
	var agency models.Agency
	result := db.Model(&models.Agency{}).First(&agency, username)

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

func DeleteAgency(agency models.Agency, db *gorm.DB) error {
	result := db.Model(&models.Agency{}).Where("username = ?", agency.Username).Delete(agency)
	return result.Error
}

func UpdateAgency(agency models.Agency, db *gorm.DB) error {
	_, err := GetAgencyByUsername(agency.Username, db)
	if err != nil {
		return err
	}
	result := db.Model(&models.Agency{}).Where("username = ?", agency.Username).Updates(agency)
	return result.Error
}
