package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterTourist godoc
// @Summary Register a tourist
// @Description Register a tourist
// @description Role allowed: everyone
// @Tags auth
// @Accept json
// @Produce json
// @Param tourist body models.TouristWithUser true "Tourist"
// @Success 200 {object} models.TouristWithUserAndToken
// @Router /auth/registerTourist [post]
func RegisterTourist(c *gin.Context) {
	tx := database.MainDB.Begin()
	var loginService LoginService = StaticLoginService()
	var jwtService JWTService = JWTAuthService()
	var loginController LoginController = LoginHandler(loginService, jwtService)
	var payload models.TouristWithUser
	if err := c.ShouldBindJSON(&payload); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	user := models.ToUserFromTouristWithUser(payload)
	user.Role = "Tourist"

	tourist := models.ToTourist(payload)

	err := database.CreateUser(&user, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	err = database.CreateTourist(&tourist, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.Params = append(c.Params, gin.Param{Key: "username", Value: user.Username})

	tokenJwt := loginController.Login(c)
	if tokenJwt == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve username from context"})
		return
	}

	touristWithUser := models.ToTouristWithUser(tourist, user)
	touristWithUserAndToken := models.ToTouristWithUserAndToken(touristWithUser, tokenJwt)

	c.JSON(http.StatusOK, gin.H{"success": true, "data": touristWithUserAndToken})
	tx.Commit()
}

// RegisterAgency godoc
// @Summary Register an agency
// @Description Register an agency
// @description Role allowed: everyone
// @Tags auth
// @Accept json
// @Produce json
// @Param agency body models.AgencyWithCompanyInformation true "Agency"
// @Success 200 {object} models.AgencyWithCompanyInformationAndToken
// @Router /auth/registerAgency [post]
func RegisterAgency(c *gin.Context) {
	tx := database.MainDB.Begin()
	var loginService LoginService = StaticLoginService()
	var jwtService JWTService = JWTAuthService()
	var loginController LoginController = LoginHandler(loginService, jwtService)
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

	err := database.CreateUser(&user, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	err = database.CreateAgency(&agency, image, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.Params = append(c.Params, gin.Param{Key: "username", Value: user.Username})

	tokenJwt := loginController.Login(c) // crap..., must move
	if tokenJwt == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve username from context"})
		return
	}
	agencyWithCompanyInformation := models.ToAgencyWithCompanyInformation(agency, user, string(image))
	agencyWithCompanyInformationAndToken := models.ToAgencyWithCompanyInformationAndToken(agencyWithCompanyInformation, tokenJwt)

	c.JSON(http.StatusOK, gin.H{"success": true, "data": agencyWithCompanyInformationAndToken})
	tx.Commit()

}
