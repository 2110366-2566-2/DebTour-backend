package models

import "time"

type Agency struct {
	Username         string     `json:"username"`
	AgencyName       string     `json:"agencyName"`
	LicenseNo        string     `json:"licenseNo"`
	BankAccount      string     `json:"bankAccount"`
	AuthorizeAdminId *int       `json:"authorizeAdminId"`
	AuthorizeStatus  string     `json:"authorizeStatus"`
	ApproveTime      *time.Time `json:"approveTime"`
}
