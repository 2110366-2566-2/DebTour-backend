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
	Phone            string     `gorm:"not null" json:"phone"`
	Email            string     `gorm:"not null" json:"email"`
	Image            string     `gorm:"not null" json:"image"`
	Role             string     `gorm:"not null" json:"role"`
	AgencyName       string     `gorm:"not null" json:"agencyName"`
	LicenseNo        string     `gorm:"unique;not null" json:"licenseNo"`
	BankAccount      string     `gorm:"not null" json:"bankAccount"`
	AuthorizeAdminId *int       `json:"authorizeAdminId"`
	AuthorizeStatus  string     `gorm:"not null" json:"authorizeStatus"`
	ApproveTime      *time.Time `json:"approveTime"`
}

type AgencyWithCompanyInformation struct {
	Username           string     `gorm:"primary_key" json:"username"`
	Phone              string     `gorm:"not null" json:"phone"`
	Email              string     `gorm:"not null" json:"email"`
	Image              string     `gorm:"not null" json:"image"`
	Role               string     `gorm:"not null" json:"role"`
	AgencyName         string     `gorm:"not null" json:"agencyName"`
	LicenseNo          string     `gorm:"unique;not null" json:"licenseNo"`
	BankAccount        string     `gorm:"not null" json:"bankAccount"`
	AuthorizeAdminId   *int       `json:"authorizeAdminId"`
	AuthorizeStatus    string     `gorm:"not null" json:"authorizeStatus"`
	ApproveTime        *time.Time `json:"approveTime"`
	CompanyInformation string     `json:"companyInformation"`
}

func ToAgencyWithUser(agency Agency, user User) AgencyWithUser {
	return AgencyWithUser{
		Username:         user.Username,
		Phone:            user.Phone,
		Email:            user.Email,
		Image:            user.Image,
		Role:             user.Role,
		AgencyName:       agency.AgencyName,
		LicenseNo:        agency.LicenseNo,
		BankAccount:      agency.BankAccount,
		AuthorizeAdminId: agency.AuthorizeAdminId,
		AuthorizeStatus:  agency.AuthorizeStatus,
		ApproveTime:      agency.ApproveTime,
	}
}

func ToAgencyWithCompanyInformation(agency Agency, user User, image string) AgencyWithCompanyInformation {
	return AgencyWithCompanyInformation{
		Username:           user.Username,
		Phone:              user.Phone,
		Email:              user.Email,
		Image:              user.Image,
		Role:               user.Role,
		AgencyName:         agency.AgencyName,
		LicenseNo:          agency.LicenseNo,
		BankAccount:        agency.BankAccount,
		AuthorizeAdminId:   agency.AuthorizeAdminId,
		AuthorizeStatus:    agency.AuthorizeStatus,
		ApproveTime:        agency.ApproveTime,
		CompanyInformation: image,
	}
}
