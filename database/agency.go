package database

import (
	"DebTour/models"
	"encoding/base64"

	"gorm.io/gorm"
)

func GetAgencyByUsername(username string, db *gorm.DB) (models.Agency, error) {
	var agency models.Agency
	result := db.Model(&models.Agency{}).Where("username = ?", username).First(&agency)

	return agency, result.Error
}

func CreateAgency(agency *models.Agency, image string, db *gorm.DB) error {
	tx := db.SavePoint("BeforeCreateAgency")

	result := tx.Model(&models.Agency{}).Create(agency)
	if result.Error != nil {
		tx.RollbackTo("BeforeCreateAgency")
		return result.Error
	}

	imageByte, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		tx.RollbackTo("BeforeCreateAgency")
		return err
	}

	companyInformation := models.CompanyInformation{Username: agency.Username, Image: imageByte}

	err = CreateCompanyInformation(&companyInformation, tx)
	if err != nil {
		tx.RollbackTo("BeforeCreateAgency")
		return err
	}

	tx.Commit()
	return nil
}

func GetAllAgencies(db *gorm.DB) ([]models.AgencyWithUser, error) {
	var agencies []models.AgencyWithUser
	result := db.Model(&models.AgencyWithUser{}).Find(&agencies)
	return agencies, result.Error
}

// create getallagencieswithcompanyinformation function
func GetAllAgenciesWithCompanyInformation(db *gorm.DB) (AgencyWithCompanyInformation []models.AgencyWithCompanyInformation, err error) {
	var agencies []models.Agency
	result := db.Model(&models.Agency{}).Find(&agencies)

	if result.Error != nil {
		return AgencyWithCompanyInformation, result.Error
	}

	var agenciesWithCompanyInformation []models.AgencyWithCompanyInformation

	for _, agency := range agencies {
		companyInformation, err := GetCompanyInformationByAgencyUsername(agency.Username, db)
		if err != nil {
			return AgencyWithCompanyInformation, err
		}
		user, err := GetUserByUsername(agency.Username, db)
		if err != nil {
			return AgencyWithCompanyInformation, err
		}
		agencyWithCompanyInformation := models.ToAgencyWithCompanyInformation(agency, user, string(companyInformation.Image))
		agenciesWithCompanyInformation = append(agenciesWithCompanyInformation, agencyWithCompanyInformation)
	}

	return agenciesWithCompanyInformation, nil
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
