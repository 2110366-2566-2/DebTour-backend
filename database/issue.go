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
			query = query.Where("reporter_username = ?", filter)
		case []string:
			query = query.Where("status = ?", filter)
		}
	}

	result := query.Find(&issues)
	return issues, result.Error
}

func GetIssueByIssueId(issueId int, db *gorm.DB) (issue models.Issue, err error) {
	result := db.Model(&models.Issue{}).Where("issue_id = ?", issueId).First(&issue)
	return issue, result.Error
}

func CreateIssue(db *gorm.DB, issue *models.Issue) error {
	result := db.Create(issue)
	return result.Error
}

func UpdateIssue(db *gorm.DB, issue *models.Issue) error {
	result := db.Model(&models.Issue{}).Where("issue_id = ?", issue.IssueId).Updates(issue)
	return result.Error
}
