package models

import "time"

type Issue struct {
	IssueId          uint       `gorm:"primaryKey" json:"issueId"`
	IssueType        string     `gorm:"not null" json:"issueType"`
	Message          string     `gorm:"not null" json:"message"`
	Status           string     `gorm:"not null" json:"status"`
	ReporterUsername string     `gorm:"not null" json:"reporterUsername"`
	ReportTimestamp  time.Time  `gorm:"autoCreateTime" json:"reportTimestamp"`
	ResolverAdminId  *int       `json:"resolverAdminId"`
	ResolveMessage   *string    `json:"resolveMessage"`
	ResolveTimestamp *time.Time `json:"resolveTimestamp"`
	Image            string     `json:"image"`
}
