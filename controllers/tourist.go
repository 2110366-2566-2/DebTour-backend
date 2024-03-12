package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//create function for create tourist

// CreateTourist godoc
// @Summary Create a tourist
// @Description Create a tourist
// @Tags tourists
// @Accept json
// @Produce json
// @Param tourist body models.Tourist true "Tourist"
// @Success 200 {object} models.Tourist
// @Router /tourists [post]
func CreateTourist(c *gin.Context) {
	var tourist models.Tourist
	if err := c.ShouldBindJSON(&tourist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	err := database.CreateTourist(&tourist, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": tourist})
}

//create function for get all tourists

// GetAllTourists godoc
// @Summary Get all tourists
// @Description Get all tourists
// @Tags tourists
// @Produce json
// @Success 200 {array} models.Tourist
// @Router /tourists [get]
func GetAllTourists(c *gin.Context) {
	tourists, err := database.GetAllTourists(database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(tourists), "data": tourists})
}

//create function for get tourist by username

// GetTouristByUsername godoc
// @Summary Get tourist by username
// @Description Get tourist by username
// @Tags tourists
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} models.Tourist
// @Router /tourists/{username} [get]
func GetTouristByUsername(c *gin.Context) {
	username := c.Param("username")
	tourist, err := database.GetTouristByUsername(username, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": tourist})
}

//create function for delete tourist

// DeleteTourist godoc
// @Summary Delete a tourist
// @Description Delete a tourist
// @Tags tourists
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} models.Tourist
// @Router /tourists/{username} [delete]
func DeleteTourist(c *gin.Context) {
	username := c.Param("username")
	tourist, err := database.GetTouristByUsername(username, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	err = database.DeleteTourist(tourist, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": tourist})
}

//create function for update tourist

// UpdateTourist godoc
// @Summary Update a tourist
// @Description Update a tourist
// @Tags tourists
// @Accept json
// @Produce json
// @Param tourist body models.Tourist true "Tourist"
// @Success 200 {object} models.Tourist
// @Router /tourists [put]
func UpdateTourist(c *gin.Context) {
	var tourist models.Tourist
	if err := c.ShouldBindJSON(&tourist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	err := database.UpdateTourist(tourist, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": tourist})
}
