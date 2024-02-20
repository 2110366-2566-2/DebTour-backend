package models

type ActivityLocation struct {
	TourId     uint `gorm:"foreignKey;not null" json:"tourId"`
	ActivityId uint `gorm:"foreignKey;not null" json:"activityId"`
	LocationId uint `gorm:"foreignKey;not null" json:"locationId"`
}