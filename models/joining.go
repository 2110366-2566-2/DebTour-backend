package models

import "gorm.io/gorm"

type Joining struct {
	TourId          uint   `gorm:"foreignKey" json:"tourId"`
	TouristUsername string `gorm:"foreignKey" json:"touristUsername"`
	MemberFirstName string `json:"memberFirstName"`
	MemberLastName  string `json:"memberLastName"`
	MemberAge       uint   `json:"memberAge"`
}

func GetALlJoinings() ([]Joining, error) {
	var joining []Joining
	result := db.Model(&Joining{}).Find(&joining)
	return joining, result.Error
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
