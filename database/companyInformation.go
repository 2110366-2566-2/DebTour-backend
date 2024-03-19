package database

import (
	"DebTour/models"
	"encoding/base64"

	"gorm.io/gorm"
)

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

func CreateCompanyInformation(companyInformation *models.CompanyInformation, db *gorm.DB) (err error) {
	db.SavePoint("BeforeCreateCompanyInformation")
	err = db.Model(&models.CompanyInformation{}).Create(&companyInformation).Error
	if err != nil {
		db.RollbackTo("BeforeCreateCompanyInformation")
		return err
	}

	return nil
}

func DeleteCompanyInformationByAgencyUsername(agencyUsername string, db *gorm.DB) (err error) {
	//check if agency exists
	_, err = GetAgencyByUsername(agencyUsername, db)
	if err != nil {
		return err
	}

	err = db.Where("username = ?", agencyUsername).Delete(&models.CompanyInformation{}).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateCompanyInformationByAgencyUsername(agencyUsername string, companyInformation models.CompanyInformation, db *gorm.DB) (err error) {
	existingCompanyInformation, err := GetCompanyInformationByAgencyUsername(agencyUsername, db)
	if err != nil {
		return err
	}
	err = db.Model(&existingCompanyInformation).Where("username = ?", agencyUsername).Updates(companyInformation).Error
	if err != nil {
		return err
	}
	return nil
}
