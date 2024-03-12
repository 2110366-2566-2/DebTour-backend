package models

import "time"

type Agency struct {
	Username         string     `gorm:"primary_key" json:"username"`
	AgencyName       string     `gorm:"not null" json:"agencyName"`
	LicenseNo        string     `gorm:"unique;not null" json:"licenseNo"`
	BankAccount      string     `gorm:"not null" json:"bankAccount"`
	AuthorizeAdminId *int       `json:"authorizeAdminId"`
	AuthorizeStatus  string     `gorm:"not null" json:"authorizeStatus"`
	ApproveTime      *time.Time `json:"approveTime"`
}

type AgencyWithUser struct {
	Username         string     `gorm:"primary_key" json:"username"`
	Password         string     `gorm:"not null" json:"password"`
	Phone            string     `gorm:"not null" json:"phone"`
	Email            string     `gorm:"not null" json:"email"`
	Image            string     `gorm:"not null" json:"image"`
	Role             string     `gorm:"not null" json:"role"`
	CreatedTime      time.Time  `gorm:"autoCreateTime"`
	AgencyName       string     `gorm:"not null" json:"agencyName"`
	LicenseNo        string     `gorm:"unique;not null" json:"licenseNo"`
	BankAccount      string     `gorm:"not null" json:"bankAccount"`
	AuthorizeAdminId *int       `json:"authorizeAdminId"`
	AuthorizeStatus  string     `gorm:"not null" json:"authorizeStatus"`
	ApproveTime      *time.Time `json:"approveTime"`
}
