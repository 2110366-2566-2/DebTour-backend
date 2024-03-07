package database

import (
	"DebTour/models"

	"gorm.io/gorm"
)

func GetAllIssues(db *gorm.DB) (issues []models.Issue, err error) {
	// find all issues in DB
	result := db.Model(&models.Issue{}).Find(&issues)

	return issues, result.Error
}

func CreateIssue(db *gorm.DB, issue *models.Issue) error {
	// Create a new issue report in DB
	result := db.Create(issue)

	return result.Error
}
