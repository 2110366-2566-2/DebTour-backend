package controllers

import (
	"DebTour/database"
	"DebTour/models"
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
// @Security ApiKeyAuth
// @produce json
// @param verifyagency body models.VerifyAgency true "VerifyAgency"
// @success 200 {object} string "Approved/Unapprved by Admin : adminusername"
// @router /agencies/verify [put]
func VerifyAgency(c *gin.Context) {
	tx := database.MainDB.Begin()
	//get username from token with bearer
	authHeader := c.GetHeader("Authorization")
	adminUsername := GetUsernameByTokenWithBearer(authHeader)
	verifyAgency := models.VerifyAgency{}
	if err := c.ShouldBindJSON(&verifyAgency); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	agencyUsername := verifyAgency.Username

	agency, err := database.GetAgencyByUsername(verifyAgency.Username, database.MainDB)
	if err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to get agency"})
		return
	}

	agency.AuthorizeAdminUsername = adminUsername
	agency.AuthorizeStatus = verifyAgency.Status

	//check if authorizestatus is "Approved" and "Unapproved" , else reject invalid format
	if agency.AuthorizeStatus != "Approved" && agency.AuthorizeStatus != "Unapproved" {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid AuthorizeStatus"})
		return
	}

	tim := time.Now()
	agency.ApproveTime = &tim

	err = database.UpdateAgencyByUsername(agencyUsername, agency, database.MainDB)
	if err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to update agency"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": agency.AuthorizeStatus + " by Admin : " + adminUsername})
}
