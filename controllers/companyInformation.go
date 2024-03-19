package controllers

import (
	"DebTour/database"
	"DebTour/models"
	//"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCompanyInformationByAgencyUsername godoc
// @Summary Get company information by agency username
// @Description Get company information by agency username
// @Tags company-informations
// @Produce json
// @Param username path string true "Agency Username"
// @Success 200 {object} models.CompanyInformationResponse "Company information"
// @Router /agencies/companyInformation/{username} [get]
func GetCompanyInformationByAgencyUsername(c *gin.Context) {
	agencyUsername := c.Param("username")

	companyInformation, err := database.GetCompanyInformationByAgencyUsername(agencyUsername, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	companyInfoResponse := models.CompanyInformationResponse{}
	companyInfoResponse.Username = agencyUsername

	c.JSON(http.StatusOK, gin.H{"success": true, "data": companyInformation})
}

// GetAllAgenciesWithCompanyInformation godoc
// @Summary Get all agencies with company information
// @Description Get all agencies with company information
// @Tags company-informations
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.AgencyWithCompanyInformation
// @Router /agencies/companyInformation [get]
func GetAllAgenciesWithCompanyInformation(c *gin.Context) {
	agencies, err := database.GetAllAgenciesWithCompanyInformation(database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(agencies), "data": agencies})
}

// DeleteCompanyInformationByAgencyUsername godoc
// @Summary Delete company information by agency username
// @Description Delete company information by agency username
// @Tags company-informations
// @Produce json
// @Param username path string true "Agency Username"
// @Success 200 {string} string "Company information deleted successfully"
// @Router /agencies/companyInformation/{username} [delete]
func DeleteCompanyInformationByAgencyUsername(c *gin.Context) {
	agencyUsername := c.Param("username")

	err := database.DeleteCompanyInformationByAgencyUsername(agencyUsername, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": "Company information deleted successfully"})
}
