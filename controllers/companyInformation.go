package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

//create function for get all company information

func GetAllCompanyInformation(c *gin.Context) {
	companyInformation, err := database.GetAllCompanyInformation(database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": companyInformation})
}

//create function for get company information by agencyusername

func GetCompanyInformationByAgencyUsername(c *gin.Context) {
	agencyUsername := c.Param("agencyusername")
	companyInformation, err := database.GetCompanyInformationByAgencyUsername(agencyUsername, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": companyInformation})
}

//create function for create company information by agency username
//use format of createtourimagebytourid

func CreateCompanyInformationByAgencyUsername(c *gin.Context) {
	var companyInformationRequest models.CompanyInformationRequest

	if err := c.ShouldBindJSON(&companyInformationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	agencyUsername := c.Param("agencyusername")

	tx := database.MainDB.Begin()

	image, err := base64.StdEncoding.DecodeString(companyInformationRequest.Image) // Convert imageb64 to string
	companyInfoImage := models.CompanyInformation{Username: agencyUsername, Image: image}
	err = database.CreateCompanyInformation(&companyInfoImage, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	tx.Commit()
	c.JSON(http.StatusCreated, gin.H{"success": true, "data": "Company information created successfully"})
}

//create function for delete company information by agency username
//use format of deletetourimagesbytourid

func DeleteCompanyInformationByAgencyUsername(c *gin.Context) {
	agencyUsername := c.Param("agencyusername")

	err := database.DeleteCompanyInformationByAgencyUsername(agencyUsername, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": "Company information deleted successfully"})
}
