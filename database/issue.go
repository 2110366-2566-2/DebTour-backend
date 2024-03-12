package database

import (
	"DebTour/models"

	"gorm.io/gorm"
)

func GetIssues(db *gorm.DB, username ...string) (issues []models.Issue, err error) {
	query := db.Model(&models.Issue{})

	if len(username) > 0 {
		query = query.Where("reporter_username = ?", username[0])
	}

	result := query.Find(&issues)
	return issues, result.Error
}

func CreateIssue(db *gorm.DB, issue *models.Issue) error {
	result := db.Create(issue)
	return result.Error
}
