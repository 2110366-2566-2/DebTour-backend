package controllers

import (
	"DebTour/database"

	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCompanyInformationByAgencyUsername godoc
// @Summary Get company information by agency username
// @Description Get company information by agency username
// @description Role allowed: "Admin" and "AgencyThemselves"
// @Tags company-information
// @Produce json
// @Security ApiKeyAuth
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

	c.JSON(http.StatusOK, gin.H{"success": true, "data": companyInformation})
}

// GetAllAgenciesWithCompanyInformation godoc
// @Summary Get all agencies with company information
// @Description Get all agencies with company information
// @description Role allowed: "Admin"
// @Tags company-information
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
// @description Role allowed: "Admin" and "AgencyThemselves"
// @Tags company-information
// @Produce json
// @Security ApiKeyAuth
// @Param username path string true "Agency Username"
// @Success 200 {string} string "Company information deleted successfully"
// @Router /agencies/companyInformation/{username} [delete]
func DeleteCompanyInformationByAgencyUsername(c *gin.Context) {
	tx := database.MainDB.Begin()
	agencyUsername := c.Param("username")

	err := database.DeleteCompanyInformationByAgencyUsername(agencyUsername, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": "Company information deleted successfully"})
}
