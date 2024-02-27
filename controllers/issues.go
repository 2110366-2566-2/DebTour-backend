package controllers

import (
	"DebTour/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllActivities godoc
// @Summary Get all activities
// @Description Get all activities
// @Tags activities
// @ID GetAllActivities
// @Produce json
// @Success 200 {array} models.Activity
// @Router /activities [get]
func GetAllIssues(c *gin.Context) {
	issues, err := database.GetAllIssues(database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": issues})
}