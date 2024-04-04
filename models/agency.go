package models

import "time"

type Agency struct {
	Username               string     `gorm:"primary_key" json:"username"`
	AgencyName             string     `gorm:"not null" json:"agencyName"`
	LicenseNo              string     `gorm:"unique;not null" json:"licenseNo"`
	BankAccount            string     `gorm:"not null" json:"bankAccount"`
	BankName               string     `gorm:"not null" json:"bankName"`
	AuthorizeAdminUsername string     `json:"authorizeAdminUsername"`
	AuthorizeStatus        string     `gorm:"not null" json:"authorizeStatus"`
	ApproveTime            *time.Time `json:"approveTime"`
	LastWithdrawTime       *time.Time `json:"lastWithdrawTime"`
}

type AgencyWithUser struct {
	Username               string     `gorm:"primary_key" json:"username"`
	Phone                  string     `gorm:"not null" json:"phone"`
	Email                  string     `gorm:"not null" json:"email"`
	Image                  string     `gorm:"not null" json:"image"`
	Role                   string     `gorm:"not null" json:"role"`
	AgencyName             string     `gorm:"not null" json:"agencyName"`
	LicenseNo              string     `gorm:"unique;not null" json:"licenseNo"`
	BankAccount            string     `gorm:"not null" json:"bankAccount"`
	BankName               string     `gorm:"not null" json:"bankName"`
	AuthorizeAdminUsername string     `json:"authorizeAdminUsername"`
	AuthorizeStatus        string     `gorm:"not null" json:"authorizeStatus"`
	ApproveTime            *time.Time `json:"approveTime"`
}

type AgencyWithCompanyInformation struct {
	Username               string     `gorm:"primary_key" json:"username"`
	Phone                  string     `gorm:"not null" json:"phone"`
	Email                  string     `gorm:"not null" json:"email"`
	Image                  string     `gorm:"not null" json:"image"`
	Role                   string     `gorm:"not null" json:"role"`
	AgencyName             string     `gorm:"not null" json:"agencyName"`
	LicenseNo              string     `gorm:"unique;not null" json:"licenseNo"`
	BankAccount            string     `gorm:"not null" json:"bankAccount"`
	BankName               string     `gorm:"not null" json:"bankName"`
	AuthorizeAdminUsername string     `json:"authorizeAdminUsername"`
	AuthorizeStatus        string     `gorm:"not null" json:"authorizeStatus"`
	ApproveTime            *time.Time `json:"approveTime"`
	CompanyInformation     string     `json:"companyInformation"`
}

type AgencyWithCompanyInformationAndToken struct {
	Username               string     `gorm:"primary_key" json:"username"`
	Phone                  string     `gorm:"not null" json:"phone"`
	Email                  string     `gorm:"not null" json:"email"`
	Image                  string     `gorm:"not null" json:"image"`
	Role                   string     `gorm:"not null" json:"role"`
	AgencyName             string     `gorm:"not null" json:"agencyName"`
	LicenseNo              string     `gorm:"unique;not null" json:"licenseNo"`
	BankAccount            string     `gorm:"not null" json:"bankAccount"`
	BankName               string     `gorm:"not null" json:"bankName"`
	AuthorizeAdminUsername string     `json:"authorizeAdminUsername"`
	AuthorizeStatus        string     `gorm:"not null" json:"authorizeStatus"`
	ApproveTime            *time.Time `json:"approveTime"`
	CompanyInformation     string     `json:"companyInformation"`
	Token                  string     `json:"token"`
}

type VerifyAgency struct {
	Username string `json:"username"`
	Status   string `json:"status"`
}

func ToAgencyWithUser(agency Agency, user User) AgencyWithUser {
	return AgencyWithUser{
		Username:               user.Username,
		Phone:                  user.Phone,
		Email:                  user.Email,
		Image:                  user.Image,
		Role:                   user.Role,
		AgencyName:             agency.AgencyName,
		LicenseNo:              agency.LicenseNo,
		BankAccount:            agency.BankAccount,
		BankName:               agency.BankName,
		AuthorizeAdminUsername: agency.AuthorizeAdminUsername,
		AuthorizeStatus:        agency.AuthorizeStatus,
		ApproveTime:            agency.ApproveTime,
	}
}

func ToAgencyWithCompanyInformation(agency Agency, user User, image string) AgencyWithCompanyInformation {
	return AgencyWithCompanyInformation{
		Username:               user.Username,
		Phone:                  user.Phone,
		Email:                  user.Email,
		Image:                  user.Image,
		Role:                   user.Role,
		AgencyName:             agency.AgencyName,
		LicenseNo:              agency.LicenseNo,
		BankAccount:            agency.BankAccount,
		BankName:               agency.BankName,
		AuthorizeAdminUsername: agency.AuthorizeAdminUsername,
		AuthorizeStatus:        agency.AuthorizeStatus,
		ApproveTime:            agency.ApproveTime,
		CompanyInformation:     image,
	}
}

func ToAgencyWithCompanyInformationAndToken(agencyWithCompanyInformation AgencyWithCompanyInformation, token string) AgencyWithCompanyInformationAndToken {
	return AgencyWithCompanyInformationAndToken{
		Username:               agencyWithCompanyInformation.Username,
		Phone:                  agencyWithCompanyInformation.Phone,
		Email:                  agencyWithCompanyInformation.Email,
		Image:                  agencyWithCompanyInformation.Image,
		Role:                   agencyWithCompanyInformation.Role,
		AgencyName:             agencyWithCompanyInformation.AgencyName,
		LicenseNo:              agencyWithCompanyInformation.LicenseNo,
		BankAccount:            agencyWithCompanyInformation.BankAccount,
		BankName:               agencyWithCompanyInformation.BankName,
		AuthorizeAdminUsername: agencyWithCompanyInformation.AuthorizeAdminUsername,
		AuthorizeStatus:        agencyWithCompanyInformation.AuthorizeStatus,
		ApproveTime:            agencyWithCompanyInformation.ApproveTime,
		CompanyInformation:     agencyWithCompanyInformation.CompanyInformation,
		Token:                  token,
	}
}

func ToAgency(agencyWithCompanyInformation AgencyWithCompanyInformation) Agency {
	return Agency{
		Username:               agencyWithCompanyInformation.Username,
		AgencyName:             agencyWithCompanyInformation.AgencyName,
		LicenseNo:              agencyWithCompanyInformation.LicenseNo,
		BankAccount:            agencyWithCompanyInformation.BankAccount,
		BankName:               agencyWithCompanyInformation.BankName,
		AuthorizeAdminUsername: agencyWithCompanyInformation.AuthorizeAdminUsername,
		AuthorizeStatus:        agencyWithCompanyInformation.AuthorizeStatus,
		ApproveTime:            agencyWithCompanyInformation.ApproveTime,
	}
}

func ToUserFromAgencyWithCompanyInformation(agencyWithCompanyInformation AgencyWithCompanyInformation) User {
	return User{
		Username: agencyWithCompanyInformation.Username,
		Phone:    agencyWithCompanyInformation.Phone,
		Email:    agencyWithCompanyInformation.Email,
		Image:    agencyWithCompanyInformation.Image,
		Role:     agencyWithCompanyInformation.Role,
	}
}
