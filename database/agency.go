package database

import (
	"DebTour/models"
	"encoding/base64"
	"errors"

	"gorm.io/gorm"
)

func GetAgencyByUsername(username string, db *gorm.DB) (agency models.Agency, err error) {
	result := db.Model(&models.Agency{}).Where("username = ?", username).First(&agency)
	return agency, result.Error
}

func GetAgencyWithUserByUsername(username string, db *gorm.DB) (agencyWithUser models.AgencyWithUser, err error) {
	var agency models.Agency
	result := db.Model(&models.Agency{}).Where("username = ?", username).First(&agency)

	user, err := GetUserByUsername(username, db)
	if err != nil {
		return agencyWithUser, err
	}
	agencyWithUser = models.ToAgencyWithUser(agency, user)

	return agencyWithUser, result.Error
}

func CreateAgency(agency *models.Agency, image string, db *gorm.DB) (err error) {
	db.SavePoint("BeforeCreateAgency")

	result := db.Model(&models.Agency{}).Create(agency)
	if result.Error != nil {
		db.RollbackTo("BeforeCreateAgency")
		return result.Error
	}

	imageByte, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		db.RollbackTo("BeforeCreateAgency")
		return err
	}

	companyInformation := models.CompanyInformation{Username: agency.Username, Image: imageByte}

	err = CreateCompanyInformation(&companyInformation, db)
	if err != nil {
		db.RollbackTo("BeforeCreateAgency")
		return err
	}

	return nil
}

func GetAllAgencies(db *gorm.DB) (agenciesWithUser []models.AgencyWithUser, err error) {
	var agencies []models.Agency
	result := db.Model(&models.Agency{}).Find(&agencies)

	// loop get user by agency.Username
	for _, agency := range agencies {
		user, err := GetUserByUsername(agency.Username, db)
		if err != nil {
			return agenciesWithUser, err
		}
		agencyWithUser := models.ToAgencyWithUser(agency, user)
		agenciesWithUser = append(agenciesWithUser, agencyWithUser)
	}

	return agenciesWithUser, result.Error
}

func GetAllAgenciesWithCompanyInformation(db *gorm.DB) (AgencyWithCompanyInformation []models.AgencyWithCompanyInformation, err error) {
	var agencies []models.Agency
	result := db.Model(&models.Agency{}).Find(&agencies)

	if result.Error != nil {
		return AgencyWithCompanyInformation, result.Error
	}

	var agenciesWithCompanyInformation []models.AgencyWithCompanyInformation

	for _, agency := range agencies {
		companyInformation, err := GetCompanyInformationByAgencyUsername(agency.Username, db)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//continue to next iteration
			continue
		}
		if err != nil {
			return AgencyWithCompanyInformation, err
		}
		user, err := GetUserByUsername(agency.Username, db)
		if err != nil {
			return AgencyWithCompanyInformation, err
		}
		agencyWithCompanyInformation := models.ToAgencyWithCompanyInformation(agency, user, companyInformation.Image)
		agenciesWithCompanyInformation = append(agenciesWithCompanyInformation, agencyWithCompanyInformation)
	}

	return agenciesWithCompanyInformation, nil
}
func DeleteAgencyByUsername(username string, db *gorm.DB) (err error) {
	//cascade delete company information
	db.SavePoint("BeforeDeleteAgency")

	err = DeleteCompanyInformationByAgencyUsername(username, db)
	if err != nil {
		db.RollbackTo("BeforeDeleteAgency")
		return err
	}

	err = db.Where("username = ?", username).Delete(&models.Agency{}).Error
	if err != nil {
		db.RollbackTo("BeforeDeleteAgency")
		return err
	}

	return err
}

func UpdateAgencyByUsername(username string, agency models.Agency, db *gorm.DB) (err error) {
	db.SavePoint("BeforeUpdateAgency")
	err = db.Model(&models.Agency{}).Where("username = ?", username).Updates(agency).Error
	if err != nil {
		db.RollbackTo("BeforeUpdateAgency")
		return err
	}

	return err
}

func GetAgencyWithCompanyInformationByUsername(username string, db *gorm.DB) (agencyWithCompanyInformation models.AgencyWithCompanyInformation, err error) {
	agency, err := GetAgencyByUsername(username, db)
	if err != nil {
		return agencyWithCompanyInformation, err
	}

	user, err := GetUserByUsername(username, db)
	if err != nil {
		return agencyWithCompanyInformation, err
	}

	companyInformation, err := GetCompanyInformationByAgencyUsername(username, db)
	if err != nil {
		return agencyWithCompanyInformation, err
	}

	agencyWithCompanyInformation = models.ToAgencyWithCompanyInformation(agency, user, companyInformation.Image)
	return agencyWithCompanyInformation, nil
}
