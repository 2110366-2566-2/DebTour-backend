package database

import (
	"DebTour/models"

	"gorm.io/gorm"
)

// get all company information
// func GetAllCompanyInformation(db *gorm.DB) ([]models.CompanyInformation, error) {
// 	var companyInformations []models.CompanyInformation

// 	err := db.Model(&models.CompanyInformation{}).Find(&companyInformations).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return companyInformations, nil
// }

// get company information by agency username
func GetCompanyInformationByAgencyUsername(agencyUsername string, db *gorm.DB) (*string, error) {
	var companyInformation models.CompanyInformation
	err := db.Model(&models.CompanyInformation{}).Where("agency_username = ?", agencyUsername).First(&companyInformation).Error
	if err != nil {
		return nil, err
	}
	var companyInfoImage string
	companyInfoImage = string(companyInformation.Image)
	return &companyInfoImage, nil
}

// create company information
func CreateCompanyInformation(companyInformation *models.CompanyInformation, db *gorm.DB) error {
	db.SavePoint("BeforeCreateCompanyInformation")

	err := db.Create(&companyInformation).Error
	if err != nil {
		db.RollbackTo("BeforeCreateCompanyInformation")
		return err
	}

	return nil
}

// delete company information by agency username
func DeleteCompanyInformationByAgencyUsername(agencyUsername string, db *gorm.DB) error {
	err := db.Where("agency_username = ?", agencyUsername).Delete(&models.CompanyInformation{}).Error
	if err != nil {
		return err
	}
	return nil
}
