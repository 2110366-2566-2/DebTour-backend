package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetAllUsers godoc
// @summary Get all users
// @description Get all users
// @tags users
// @id GetAllUsers
// @produce json
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
// @tags users
// @id GetUserByUsername
// @produce json
// @param username path string true "Username"
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

// DeleteUser godoc
// @summary Delete a user
// @description Delete a user
// @tags users
// @id DeleteUser
// @param username path string true "Username"
// @produce json
// @response 200 {string} string "User deleted"
// @router /users/{username} [delete]
func DeleteUserByUsername(c *gin.Context) {
	// Extract the username from the URL path parameters
	username := c.Param("username")
	//check if user exist
	_, err := database.GetUserByUsername(username, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Invalid username"})
		return
	}
	// Delete the user by username
	err = database.DeleteUserByUsername(username, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": "User deleted successfully"})
}

// UpdateUser godoc
// @summary Update a user
// @description Update a user
// @tags users
// @id UpdateUser
// @accept json
// @produce json
// @param username path string true "Username"
// @param user body models.User true "User"
// @response 200 {string} string "User updated"
// @router /users/{username} [put]
// func UpdateUser(c *gin.Context) {
// 	var user models.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
// 		return
// 	}
// 	err := database.UpdateUser(&user, database.MainDB)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"success": true, "data": user})
// }

func UpdateUserByUsername(c *gin.Context) {
	// Extract the username from the URL path parameters
	username := c.Param("username")

	// Bind the request body to a user struct
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	//get user by username check if username exists
	_, err := database.GetUserByUsername(username, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Invalid username"})
		return
	}

	// Call the database function to update the user
	err = database.UpdateUserByUsername(username, user, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": "User updated successfully"})
}

// GetMe godoc
// @Summary Get user info
// @Description Get user info
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.TouristWithUser
// @Success 200 {object} models.AgencyWithUser
// @Router /getMe [get]
func GetMe(c *gin.Context) {
	//decode token and get username
	//token will be passed in header
	tokenString := c.GetHeader("Authorization")
	token, err := JWTAuthService().ValidateToken(tokenString)
	if !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "error": err.Error()})
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	user, err := database.GetUserByUsername(username, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	if user.Role == "Tourist" {
		var data models.TouristWithUser
		tourist, err := database.GetTouristByUsername(username, database.MainDB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
		data.Username = user.Username
		data.Phone = user.Phone
		data.Email = user.Email
		data.Image = user.Image
		data.Role = user.Role
		data.CitizenId = tourist.CitizenId
		data.FirstName = tourist.FirstName
		data.LastName = tourist.LastName
		data.Address = tourist.Address
		data.BirthDate = tourist.BirthDate
		c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
	}

	if user.Role == "Agency" {
		var data models.AgencyWithUser
		agency, err := database.GetAgencyByUsername(username, database.MainDB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
		data.Username = user.Username
		data.Phone = user.Phone
		data.Email = user.Email
		data.Image = user.Image
		data.Role = user.Role
		data.AgencyName = agency.AgencyName
		data.LicenseNo = agency.LicenseNo
		data.BankAccount = agency.BankAccount
		data.AuthorizeAdminId = agency.AuthorizeAdminId
		data.AuthorizeStatus = agency.AuthorizeStatus
		data.ApproveTime = agency.ApproveTime
		c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
	}
	c.JSON(http.StatusOK, gin.H{"success": false, "error": "record not found"})
}
