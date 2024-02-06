package models

type TourImage struct {
	TourImage uint   `gorm:"primaryKey" json:"tourImage"`
	Image     string `gorm:"primaryKey" json:"image"`
}
