package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetIssues godoc
// @Summary Get issues
// @Description Get issues optionally filtered by username and/or status
// @description Role allowed: "Admin", "AgencyOwner" and "TouristOwner"
// @Tags issues
// @ID GetIssues
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.Issue
// @Router /issues [get]
func GetIssues(c *gin.Context) {
	username := c.Query("username")
	status := c.Query("status")

	var issues []models.Issue
	var err error

	if username == "" && status == "" {
		issues, err = database.GetIssues(database.MainDB)
	} else if username != "" && status != "" {
		issues, err = database.GetIssues(database.MainDB, username, []string{status})
	} else {
		issues, err = database.GetIssues(database.MainDB, username)
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
// @description Role allowed: "Admin", "Agency" and "Tourist"
// @Tags issues
// @Accept json
// @Produce json
// @Param Issue body models.Issue true "Issue object"
// @Security ApiKeyAuth
// @Success 200 {object} models.Issue
// @Router /issues [post]
func CreateIssueReport(c *gin.Context) {
	var issue models.Issue
	if err := c.ShouldBindJSON(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.MainDB.Create(&issue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": issue})

}

// UpdateIssueReport godoc
// @Summary Update an issue report
// @Description Update an existing issue report
// @description Role allowed: "Admin"
// @Tags issues
// @Accept json
// @Produce json
// @Param issue_id path string true "Issue ID"
// @Param issue body models.Issue true "Issue object"
// @Security ApiKeyAuth
// @Success 200 {object} models.Issue
// @Router /issues/{issue_id} [put]
func UpdateIssueReport(c *gin.Context) {
	issueID, err := strconv.Atoi(c.Param("issue_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid issue ID"})
		return
	}

	// Fetch the issue
	var issue models.Issue
	if err := database.MainDB.First(&issue, issueID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	// Update the issue
	if err := c.ShouldBindJSON(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database.UpdateIssue(database.MainDB, &issue)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": issue})
}
