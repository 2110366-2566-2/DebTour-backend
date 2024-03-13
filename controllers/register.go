package controllers

import (
	"DebTour/database"
	"DebTour/models"

	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHandler godoc
// @Summary Register a user
// @Description Register a user
// @Tags register
// @Accept  json
// @Produce  json
// @Param role path string true "Role"
// @Param touristwithuser body models.TouristWithUser true "TouristWithUser"
// @Param agencywithuser body models.AgencyWithUser true "AgencyWithUser"
// @Success 200 {object} models.User
// @Router /register/{role} [post]
func RegisterHandler(c *gin.Context) {
	role := c.Param("role")
	if role == "tourist" {
		RegisterTourist(c)
		return
	}
	if role == "agency" {
		RegisterAgency(c)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid role"})
}

func RegisterTourist(c *gin.Context) {
	tx := database.MainDB.Begin()
	var loginService LoginService = StaticLoginService()
	var jwtService JWTService = JWTAuthService()
	var loginController LoginController = LoginHandler(loginService, jwtService)
	role := c.Param("role")
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
	user.Role = role

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

	c.Params = append(c.Params, gin.Param{Key: "username", Value: user.Username})
	// c.Params = append(c.Params, gin.Param{Key: "role", Value: Role})

	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>> Params: ", c.Params)
	token_jwt := loginController.Login(c) // crap..., must move
	if token_jwt == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve username from context"})
		return
	}

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
		"token":          token_jwt,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
	tx.Commit()
}

func RegisterAgency(c *gin.Context) {
	tx := database.MainDB.Begin()
	var loginService LoginService = StaticLoginService()
	var jwtService JWTService = JWTAuthService()
	var loginController LoginController = LoginHandler(loginService, jwtService)
	role := c.Param("role")
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
	user.Role = role

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

	c.Params = append(c.Params, gin.Param{Key: "username", Value: user.Username})
	// c.Params = append(c.Params, gin.Param{Key: "role", Value: Role})

	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>> Params: ", c.Params)
	token_jwt := loginController.Login(c) // crap..., must move
	if token_jwt == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve username from context"})
		return
	}

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
		"token":            token_jwt,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
	tx.Commit()

}
