package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllAgencies godoc
// @Summary Get all agencies
// @Description Get all agencies
// @Tags agencies
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.AgencyWithUser
// @Router /agencies [get]
func GetAllAgencies(c *gin.Context) {
	agencies, err := database.GetAllAgencies(database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(agencies), "data": agencies})
}

// GetAgencyWithUserByUsername godoc
// @Summary Get agency with user by username
// @Description Get agency with user by username
// @Tags agencies
// @Produce json
// @Param username path string true "Username"
// @Security ApiKeyAuth
// @Success 200 {object} models.AgencyWithUser
// @Router /agencies/{username} [get]
func GetAgencyWithUserByUsername(c *gin.Context) {
	username := c.Param("username")
	agencyWithUser, err := database.GetAgencyWithUserByUsername(username, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": agencyWithUser})
}

// UpdateAgency godoc
// @Summary Update a agency
// @Description Update a agency
// @Tags agencies
// @Accept json
// @Produce json
// @Param agency body models.AgencyWithCompanyInformation true "Agency"
// @Security ApiKeyAuth
// @Success 200 {object} models.AgencyWithCompanyInformation
// @Router /agencies [put]
func UpdateAgency(c *gin.Context) {
	tx := database.MainDB.Begin()
	username := c.Param("username")
	var payload models.AgencyWithCompanyInformation
	if err := c.ShouldBindJSON(&payload); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	user := models.ToUserFromAgencyWithCompanyInformation(payload)
	user.Role = "Agency"

	agency := models.ToAgency(payload)

	image := payload.CompanyInformation

	// Now you can access agencyWithUser.User and agencyWithUser.Agency
	err := database.UpdateUserByUsername(username, user, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	err = database.UpdateAgencyByUsername(username, agency, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Create combined data
	agencyWithCompanyInformation := models.ToAgencyWithCompanyInformation(agency, user, image)

	c.JSON(http.StatusOK, gin.H{"success": true, "data": agencyWithCompanyInformation})
	tx.Commit()

}

// DeleteAgency godoc
// @Summary Delete a agency
// @Description Delete a agency
// @Tags agencies
// @Produce json
// @Param username path string true "Username"
// @Security ApiKeyAuth
// @Success 200 {string} string "Agency deleted successfully"
// @Router /agencies/{username} [delete]
func DeleteAgency(c *gin.Context) {
	tx := database.MainDB.Begin()
	username := c.Param("username")
	//check if user exist
	_, err := database.GetUserByUsername(username, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	err = database.DeleteUserByUsername(username, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	err = database.DeleteAgencyByUsername(username, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	err = database.DeleteCompanyInformationByAgencyUsername(username, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": "Agency deleted successfully"})
	tx.Commit()
}
