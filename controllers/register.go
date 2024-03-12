package controllers

import (
	"DebTour/database"
	"DebTour/models"

	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSON payload must be in the form of: TouristWithUser

// RegisterTourist godoc
// @Summary Register a tourist
// @Description Register a tourist and create a user
// @Tags tourist
// @Accept  json
// @Produce  json
// @Param tourist body TouristWithUser true "TouristWithUser"
// @Success 200 {object} models.TouristWithUser
func RegisterTourist(c *gin.Context) {
	tx := database.MainDB.Begin()
	var payload models.TouristWithUser
	if err := c.ShouldBindJSON(&payload); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	var user models.User
	// Populate user fields from payload
	user.Username = payload.Username
	user.Password = payload.Password
	user.Phone = payload.Phone
	user.Email = payload.Email
	user.Image = payload.Image
	user.Role = payload.Role

	var tourist models.Tourist
	tourist.Username = payload.Username
	tourist.CitizenId = payload.CitizenId
	tourist.FirstName = payload.FirstName
	tourist.LastName = payload.LastName
	tourist.Address = payload.Address
	tourist.BirthDate = payload.BirthDate
	tourist.Gender = payload.Gender
	tourist.DefaultPayment = payload.DefaultPayment

	// Now you can access touristWithUser.User and touristWithUser.Tourist
	err := database.CreateUser(&user, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	//print data in touristWithUser

	err = database.CreateTourist(&tourist, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	// Create combined data
	data := gin.H{
		"username":       user.Username,
		"password":       user.Password,
		"phone":          user.Phone,
		"email":          user.Email,
		"image":          user.Image,
		"role":           user.Role,
		"created_time":   user.CreatedTime,
		"citizenId":      tourist.CitizenId,
		"firstName":      tourist.FirstName,
		"lastName":       tourist.LastName,
		"address":        tourist.Address,
		"birthDate":      tourist.BirthDate,
		"gender":         tourist.Gender,
		"defaultPayment": tourist.DefaultPayment,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
	tx.Commit()
}

// RegisterAgency godoc
// @Summary Register an agency
// @Description Register an agency and create a user
// @Tags agency
// @Accept  json
// @Produce  json
// @Param agency body AgencyWithUser true "AgencyWithUser"
// @Success 200 {object} models.AgencyWithUser
func RegisterAgency(c *gin.Context) {
	tx := database.MainDB.Begin()
	var payload models.AgencyWithUser
	if err := c.ShouldBindJSON(&payload); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	var user models.User
	// Populate user fields from payload
	user.Username = payload.Username
	user.Password = payload.Password
	user.Phone = payload.Phone
	user.Email = payload.Email
	user.Image = payload.Image
	user.Role = payload.Role

	var agency models.Agency
	agency.Username = payload.Username
	agency.AgencyName = payload.AgencyName
	agency.LicenseNo = payload.LicenseNo
	agency.BankAccount = payload.BankAccount
	agency.AuthorizeAdminId = payload.AuthorizeAdminId
	agency.AuthorizeStatus = payload.AuthorizeStatus
	agency.ApproveTime = payload.ApproveTime

	// Now you can access agencyWithUser.User and agencyWithUser.Agency
	err := database.CreateUser(&user, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	err = database.CreateAgency(&agency, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Create combined data
	data := gin.H{
		"username":         user.Username,
		"password":         user.Password,
		"phone":            user.Phone,
		"email":            user.Email,
		"image":            user.Image,
		"role":             user.Role,
		"created_time":     user.CreatedTime,
		"agencyName":       agency.AgencyName,
		"licenseNo":        agency.LicenseNo,
		"bankAccount":      agency.BankAccount,
		"authorizeAdminId": agency.AuthorizeAdminId,
		"authorizeStatus":  agency.AuthorizeStatus,
		"approveTime":      agency.ApproveTime,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
	tx.Commit()

}
