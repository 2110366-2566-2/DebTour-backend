package models

import "gorm.io/gorm"

type Joining struct {
	TourId          uint   `gorm:"primaryKey" json:"tourId"`
	TouristUsername string `gorm:"primaryKey" json:"touristUsername"`
	MemberName      string `gorm:"primaryKey" json:"memberName"`
	Age             uint   `json:"age"`
}

func GetJoiningByTourId(tourId uint) (joinings []Joining, err error) {
	result := db.Model(&Joining{}).Where("tour_id = ?", tourId).Find(&joinings)

	return joinings, result.Error
}

func CreateJoining(joining *Joining) (err error) {
	result := db.Model(&Joining{}).Create(joining)

	return result.Error
}

func DeleteJoiningByTourId(tourId uint, db *gorm.DB) (err error) {
	result := db.Model(&Joining{}).Where("tour_id = ?", tourId).Delete(&Joining{})

	return result.Error
}
