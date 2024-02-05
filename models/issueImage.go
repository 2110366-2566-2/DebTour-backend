package models

type IssueImage struct {
	IssueId uint   `gorm:"primaryKey" json:"issue_id"`
	Image   string `gorm:"primaryKey" json:"image"`
}
