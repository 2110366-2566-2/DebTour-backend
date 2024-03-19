package controllers

import (
	"DebTour/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ApproveAgency godoc
// @summary Approve agency
// @description Approve agency
// @description Role allowed: "Admin"
// @tags admin
// @id ApproveAgency
// @param username path string true "Agency Username"
// @Security ApiKeyAuth
// @produce json
// @success 200 {object} models.AgencyWithCompanyInformation
// @router /agencies/verify/{username} [put]
func ApproveAgency(c *gin.Context) {
	tx := database.MainDB.Begin()
	agencyUsername := c.Param("username")
	tokenS := c.GetHeader("Authorization")
	const BEARER_SCHEMA = "Bearer "
	tokenS = tokenS[len(BEARER_SCHEMA):]
	adminUsername := GetUsernameByToken(tokenS)
	agency, err := database.GetAgencyByUsername(agencyUsername, database.MainDB)
	if err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Agency not found"})
		return
	}
	agency.AuthorizeStatus = "Approved"
	agency.AuthorizeAdminUsername = adminUsername
	tim := time.Now()
	agency.ApproveTime = &tim
	err = database.UpdateAgencyByUsername(agencyUsername, agency, database.MainDB)
	if err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update agency"})
		return
	}

	//create agencyWithCompanyInformation by agency username
	agencyWithCompanyInformation, err := database.GetAgencyWithCompanyInformationByUsername(agencyUsername, database.MainDB)
	if err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get company information"})
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": agencyWithCompanyInformation})
}
