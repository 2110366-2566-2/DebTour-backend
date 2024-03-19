package database

import (
	"DebTour/models"
	"encoding/base64"
	"fmt"

	"gorm.io/gorm"
)

func GetAgencyByUsername(username string, db *gorm.DB) (models.Agency, error) {
	var agency models.Agency
	result := db.Model(&models.Agency{}).Where("username = ?", username).First(&agency)

	return agency, result.Error
}

func CreateAgency(agency *models.Agency, image string, db *gorm.DB) error {
	tx := db.SavePoint("BeforeCreateAgency")
	fmt.Println(">>>>>>>>>>>>>", agency.BankName)
	result := tx.Model(&models.Agency{}).Create(agency)
	if result.Error != nil {
		tx.RollbackTo("BeforeCreateAgency")
		return result.Error
	}
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>> Before decode")
	imageByte, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		tx.RollbackTo("BeforeCreateAgency")
		return err
	}
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>> After decode")
	companyInformation := models.CompanyInformation{Username: agency.Username, Image: imageByte}

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>> Before CreateCompanyInformation")
	err = CreateCompanyInformation(&companyInformation, tx)
	if err != nil {
		tx.RollbackTo("BeforeCreateAgency")
		return err
	}
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>> After CreateCompanyInformation")
	tx.Commit()
	return nil
}

func GetAllAgencies(db *gorm.DB) (agenciesWithUser []models.AgencyWithUser, err error) {
	var agencies []models.Agency
	result := db.Model(&models.Agency{}).Find(&agencies)
	//loop get user by agency.Username

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
		if err == gorm.ErrRecordNotFound {
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
