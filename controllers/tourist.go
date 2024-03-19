package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllTourists godoc
// @Summary Get all tourists
// @Description Get all tourists
// @Tags tourists
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.TouristWithUser
// @Router /tourists [get]
func GetAllTouristsWithUser(c *gin.Context) {
	touristsWithUser, err := database.GetAllTourists(database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(touristsWithUser), "data": touristsWithUser})
}

// GetTouristByUsername godoc
// @Summary Get tourist by username
// @Description Get tourist by username
// @Tags tourists
// @Produce json
// @Param username path string true "Username"
// @Security ApiKeyAuth
// @Success 200 {object} models.TouristWithUser
// @Router /tourists/{username} [get]
func GetTouristByUsername(c *gin.Context) {
	username := c.Param("username")
	touristsWithUser, err := database.GetTouristByUsername(username, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": touristsWithUser})
}

// DeleteTouristByUsername godoc
// @Summary Delete tourist and user
// @Description Delete tourist and user by username
// @Tags tourists
// @Produce json
// @Param username path string true "Username"
// @Security ApiKeyAuth
// @Success 200 {string} string	"Tourist deleted successfully"
// @Router /tourists/{username} [delete]
func DeleteTouristByUsername(c *gin.Context) {
	tx := database.MainDB.Begin()
	username := c.Param("username")

	//check is username exist
	_, err := database.GetUserByUsername(username, database.MainDB)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	err = database.DeleteUserByUsername(username, database.MainDB)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	err = database.DeleteTouristByUsername(username, database.MainDB)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": "Tourist deleted successfully"})
	tx.Commit()
}

// UpdateTouristByUsername godoc
// @Summary Update a tourist
// @Description Update a tourist and user also
// @Tags tourists
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Param tourist body models.TouristWithUser true "Tourist"
// @Security ApiKeyAuth
// @Success 200 {object} models.TouristWithUser
// @Router /tourists/{username} [put]
func UpdateTouristByUsername(c *gin.Context) {
	tx := database.MainDB.Begin()
	username := c.Param("username")
	var payload models.TouristWithUser
	if err := c.ShouldBindJSON(&payload); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	user := models.ToUserFromTouristWithUser(payload)
	user.Role = "Tourist"

	tourist := models.ToTourist(payload)

	err := database.UpdateUserByUsername(username, user, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	err = database.UpdateTouristByUsername(username, tourist, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	touristWithUser := models.ToTouristWithUser(tourist, user)

	c.JSON(http.StatusOK, gin.H{"success": true, "data": touristWithUser})
	tx.Commit()
}
