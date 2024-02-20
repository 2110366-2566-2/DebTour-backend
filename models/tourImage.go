package models

type TourImage struct {
	TourId uint   `gorm:"foreignKey;not null" json:"tourId"`
	Image     string `gorm:"primaryKey" json:"image"`
}
