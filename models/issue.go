package models

import "time"

type Issue struct {
	IssueId          uint       `gorm:"primaryKey" json:"issueId"`
	IssueType        string     `json:"issueType"`
	Message          string     `json:"message"`
	Status           string     `json:"status"`
	ReporterUsername string     `json:"reporterUsername"`
	ReportTimestamp  time.Time  `gorm:"autoCreateTime" json:"reportTimestamp"`
	ResolverAdminId  *int       `json:"resolverAdminId"`
	ResolveMessage   *string    `json:"resolveMessage"`
	ResolveTimestamp *time.Time `json:"resolveTimestamp"`
}
