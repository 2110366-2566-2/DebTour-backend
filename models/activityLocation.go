package models

type ActivityLocation struct {
	TourId     uint `gorm:"primaryKey" json:"tourId"`
	ActivityId uint `gorm:"primaryKey" json:"activityId"`
	LocationId uint `gorm:"primaryKey" json:"locationId"`
}
