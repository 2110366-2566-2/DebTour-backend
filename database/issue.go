package database

import (
	"DebTour/models"

	"gorm.io/gorm"
)

func GetAllIssues(db *gorm.DB) (issues []models.Issue, err error) {
	// find all issues in the database
	result := db.Model(&models.Issue{}).Find(&issues)

	return issues, result.Error
}
