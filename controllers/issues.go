package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetIssues godoc
// @Summary Get issues
// @Description Get issues optionally filtered by username
// @Tags issues
// @ID GetIssues
// @Produce json
// @Param username query string false "Username to filter issues"
// @Success 200 {array} models.Issue
// @Router /issues [get]
func GetIssues(c *gin.Context) {
	username := c.Query("username")

	var issues []models.Issue
	var err error

	if username != "" {
		issues, err = database.GetIssues(database.MainDB, username)
	} else {
		issues, err = database.GetIssues(database.MainDB)
	}

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
