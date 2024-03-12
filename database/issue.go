package database

import (
	"DebTour/models"

	"gorm.io/gorm"
)

func GetIssues(db *gorm.DB, filters ...interface{}) (issues []models.Issue, err error) {
	query := db.Model(&models.Issue{})

	for _, filter := range filters {
		switch filter := filter.(type) {
		case string:
			query = query.Where("reporterUsername = ?", filter)
		case []string:
			query = query.Where("status IN ?", filter)
		}
	}

	result := query.Find(&issues)
	return issues, result.Error
}

func CreateIssue(db *gorm.DB, issue *models.Issue) error {
	result := db.Create(issue)
	return result.Error
}
