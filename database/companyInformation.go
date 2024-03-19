package database

import (
	"DebTour/models"
	"encoding/base64"

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
func GetCompanyInformationByAgencyUsername(agencyUsername string, db *gorm.DB) (companyInformationResponse models.CompanyInformationResponse, err error) {
	var companyInformation models.CompanyInformation
	err = db.Model(&models.CompanyInformation{}).Where("username = ?", agencyUsername).First(&companyInformation).Error
	if err != nil {
		return companyInformationResponse, err
	}

	var companyInfoImage = base64.StdEncoding.EncodeToString(companyInformation.Image)
	companyInformationResponse = models.CompanyInformationResponse{Username: agencyUsername, Image: companyInfoImage}
	return companyInformationResponse, nil
}

// create company information
func CreateCompanyInformation(companyInformation *models.CompanyInformation, db *gorm.DB) (err error) {
	db.SavePoint("BeforeCreateCompanyInformation")
	err = db.Model(&models.CompanyInformation{}).Create(&companyInformation).Error
	if err != nil {
		db.RollbackTo("BeforeCreateCompanyInformation")
		return err
	}

	return nil
}

// delete company information by agency username
func DeleteCompanyInformationByAgencyUsername(agencyUsername string, db *gorm.DB) (err error) {
	err = db.Where("username = ?", agencyUsername).Delete(&models.CompanyInformation{}).Error
	if err != nil {
		return err
	}
	return nil
}
