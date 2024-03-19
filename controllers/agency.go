package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//create function for get all agencies and company information

// GetAllAgencies godoc
// @Summary Get all agencies
// @Description Get all agencies
// @Tags agencies
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.Agency
// @Router /agencies [get]
func GetAllAgencies(c *gin.Context) {
	agencies, err := database.GetAllAgencies(database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(agencies), "data": agencies})
}

//create getallagencieswithcompanyinformation function



//create function for get agency by username

// GetAgencyByUsername godoc
// @Summary Get agency by username
// @Description Get agency by username
// @Tags agencies
// @Produce json
// @Param username path string true "Username"
// @Security ApiKeyAuth
// @Success 200 {object} models.Agency
// @Router /agencies/{username} [get]
func GetAgencyByUsername(c *gin.Context) {
	username := c.Param("username")
	agency, err := database.GetAgencyByUsername(username, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": agency})
}

//create function for update agency

// UpdateAgency godoc
// @Summary Update a agency
// @Description Update a agency
// @Tags agencies
// @Accept json
// @Produce json
// @Param agency body models.Agency true "Agency"
// @Security ApiKeyAuth
// @Success 200 {object} models.Agency
// @Router /agencies [put]
func UpdateAgency(c *gin.Context) {
	tx := database.MainDB.Begin()
	username := c.Param("username")
	var payload models.AgencyWithUser
	if err := c.ShouldBindJSON(&payload); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	var user models.User
	// Populate user fields from payload
	user.Username = payload.Username
	user.Phone = payload.Phone
	user.Email = payload.Email
	user.Image = payload.Image
	user.Role = "Agency"

	var agency models.Agency
	agency.Username = payload.Username
	agency.AgencyName = payload.AgencyName
	agency.LicenseNo = payload.LicenseNo
	agency.BankAccount = payload.BankAccount
	agency.BankName = payload.BankName
	agency.AuthorizeAdminUsername = payload.AuthorizeAdminUsername
	agency.AuthorizeStatus = payload.AuthorizeStatus
	agency.ApproveTime = payload.ApproveTime

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
	agencyWithUser := models.ToAgencyWithUser(agency, user)

	c.JSON(http.StatusOK, gin.H{"success": true, "data": agencyWithUser})
	tx.Commit()

}

//create function for delete agency

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

	c.JSON(http.StatusOK, gin.H{"success": true, "data": "Agency deleted successfully"})
	tx.Commit()
}
