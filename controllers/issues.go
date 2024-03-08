package controllers

import (
	"DebTour/database"
	"DebTour/models"
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

// CreateIssueReport godoc
// @Summary Create an issue report
// @Description Create a new issue report
// @Tags issues
// @Accept json
// @Produce json
// @Success 200 {object} models.Issue
// @Router /issues [post]
func CreateIssueReport(c *gin.Context) {
	// Create a new issue obj
	var issue models.Issue
	if err := c.ShouldBindJSON(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save an issue to DB
	if err := database.MainDB.Create(&issue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": issue})

}
