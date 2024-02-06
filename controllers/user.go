package controllers

import (
	"DebTour/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllUsers godoc
// @summary Get all users
// @description Get all users
// @id GetAllUsers
// @produce json
// @response 200 {array} User
// @router /users [get]
func GetAllUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUserByUsername godoc
// @summary Get user by username
// @description Get user by username
// @id GetUserByUsername
// @produce json
// @param username path string true "Username"
// @response 200 {object} User
// @router /users/{username} [get]
func GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := models.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// CreateUser godoc
// @summary Create a user
// @description Create a user
// @id CreateUser
// @accept json
// @produce json
// @param user body User true "User"
// @success 200 {object} User
// @router /users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := models.CreateUser(user.Username, user.Password, user.Phone, user.Email, user.Image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// DeleteUser godoc
// @summary Delete a user
// @description Delete a user
// @id DeleteUser
// @param username path string true "Username"
// @produce json
// @response 200 {string} string "User deleted"
// @router /users/{username} [delete]
func DeleteUser(c *gin.Context) {
	username := c.Param("username")
	err := models.DeleteUser(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "User deleted")
}

// UpdateUser godoc
// @summary Update a user
// @description Update a user
// @id UpdateUser
// @accept json
// @produce json
// @param username path string true "Username"
// @param user body User true "User"
// @response 200 {string} string "User updated"
// @router /users/{username} [put]
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	username := user.Username
	err := models.UpdateUser(username, user.Password, user.Phone, user.Email, user.Image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "User updated")
}