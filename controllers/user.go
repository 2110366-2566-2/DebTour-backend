package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"net/http"

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
func CreateUser(c *gin.Context) {
	var user models.User
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
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	err := database.DeleteUser(&user, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": "Delete user successfully"})
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
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	err := database.UpdateUser(&user, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": user})
}
