package models

type CompanyInformation struct {
	Username string `gorm:"foreignKey;not null" json:"username"`
	Image    []byte `gorm:"type:bytea;not null" json:"image"`
}

type CompanyInformationRequest struct {
	Image string `json:"image"`
}

// type CompanyInformationResponse struct {
// 	Username string   `json:"username"`
// 	Images   []string `json:"images"`
// }
