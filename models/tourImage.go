package models

type TourImage struct {
	TourId uint   `gorm:"foreignKey;not null" json:"tourId"`
	Image	[]byte `gorm:"type:bytea;not null" json:"image"`
}

type TourImagesRequest struct {
	Images []string `json:"images"`
}

type TourImagesResponse struct {
	TourId uint     `json:"tourId"`
	Images []string `json:"images"`
}