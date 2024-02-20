package models

type IssueImage struct {
	IssueId uint   `gorm:"foreignKey;not null" json:"issue_id"`
	Image   string `gorm:"primaryKey" json:"image"`
}
