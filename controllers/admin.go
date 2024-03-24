package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// VerifyAgency godoc
// @summary Verify agency
// @description Verify agency by admin by specify the status
// @description Role allowed: "Admin"
// @tags admin
// @id VerifyAgency
// @Security ApiKeyAuth
// @produce json
// @param VerifyAgency body models.VerifyAgency true "VerifyAgency"
// @success 200 {object} string "Approved/Unapproved by Admin : adminUsername"
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
	agency, err := database.GetAgencyByUsername(verifyAgency.Username, tx)
	if err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to get agency"})
		return
	}

	agency.AuthorizeAdminUsername = adminUsername
	agency.AuthorizeStatus = verifyAgency.Status

	//check if AuthorizeStatus is "Approved" and "Unapproved" , else reject invalid format
	if agency.AuthorizeStatus != "Approved" && agency.AuthorizeStatus != "Unapproved" {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid AuthorizeStatus"})
		return
	}

	tim := time.Now()
	agency.ApproveTime = &tim

	err = database.UpdateAgencyByUsername(agencyUsername, agency, tx)
	if err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to update agency"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": agency.AuthorizeStatus + " by Admin : " + adminUsername})
}
