package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllUsers godoc
// @summary Get all users
// @description Get all users
// @description Role allowed: "Admin"
// @tags users
// @id GetAllUsers
// @produce json
// @Security ApiKeyAuth
// @response 200 {array} models.User
// @router /users [get]
func GetAllUsers(c *gin.Context) {
	users, err := database.GetAllUsers(database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(users), "data": users})
}

// GetUserByUsername godoc
// @summary Get user by username
// @description Get user by username
// @description Role allowed: "Admin"
// @tags users
// @id GetUserByUsername
// @produce json
// @param username path string true "Username"
// @Security ApiKeyAuth
// @response 200 {object} models.User
// @router /users/{username} [get]
func GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := database.GetUserByUsername(username, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": user})
}

// CreateUser godoc
// @summary Create a user
// @description Create a user
// @tags users
// @id CreateUser
// @accept json
// @produce json
// @param user body models.User true "User"
// @success 200 {object} models.User
// @router /users [post]
type CreateUserInput struct {
	Username string `json:"username"`
}

func CreateUser(c *gin.Context) {
	var user models.User
	var username CreateUserInput
	c.ShouldBindJSON(&username)
	fmt.Println(">>>>>>>>>>>>>>>> User: ", username.Username)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	err := database.CreateUser(&user, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": user})
}

// DeleteUserByUsername godoc
// @summary Delete user by username
// @description Delete user by username
// @description Role allowed: "Admin"
// @tags users
// @id DeleteUserByUsername
// @param username path string true "Username"
// @produce json
// @Security ApiKeyAuth
// @success 200 {string} string "User deleted successfully"
// @router /users/{username} [delete]
func DeleteUserByUsername(c *gin.Context) {
	tx := database.MainDB.Begin()
	// Extract the username from the URL path parameters
	username := c.Param("username")
	//check if user exist
	_, err := database.GetUserByUsername(username, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Invalid username"})
		return
	}
	// Delete the user by username in Tourist table
	fmt.Println("Before delete tourist")
	err = database.DeleteTouristByUsername(username, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	fmt.Println("After delete tourist")

	// Delete company information by username in CompanyInformation table
	err = database.DeleteCompanyInformationByAgencyUsername(username, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Delete the user by username in Agency table
	fmt.Println("Before delete Agency")

	err = database.DeleteAgencyByUsername(username, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	fmt.Println("After delete Agency")

	// Delete the user by username in User table
	err = database.DeleteUserByUsername(username, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": "User deleted successfully"})
}

// GetMe godoc
// @Summary Get user info
// @Description Get user info
// @description Role allowed: "Admin", "Agency" and "Tourist"
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.TouristWithUser
// @Success 200 {object} models.AgencyWithUser
// @Router /getMe [get]
func GetMe(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	username := GetUsernameByTokenWithBearer(authHeader)

	user, err := database.GetUserByUsername(username, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	if user.Role == "Tourist" {
		touristWithUser, err := database.GetTouristByUsername(username, database.MainDB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "data": touristWithUser})
		return
	}

	if user.Role == "Agency" {
		agencyWithUser, err := database.GetAgencyWithUserByUsername(username, database.MainDB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "data": agencyWithUser})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": false, "error": "record not found"})
}
