package models

import "time"

type Agency struct {
	Username         string     `gorm:"primary_key" json:"username"`
	AgencyName       string     `json:"agencyName"`
	LicenseNo        string     `gorm:"unique" json:"licenseNo"`
	BankAccount      string     `json:"bankAccount"`
	AuthorizeAdminId *int       `json:"authorizeAdminId"`
	AuthorizeStatus  string     `json:"authorizeStatus"`
	ApproveTime      *time.Time `json:"approveTime"`
}
