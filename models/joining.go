package models

import (
	"gorm.io/gorm"
)

type Joining struct {
	TourId          uint   `gorm:"foreignKey" json:"tourId"`
	TouristUsername string `gorm:"foreignKey" json:"touristUsername"`
	MemberFirstName string `json:"memberFirstName"`
	MemberLastName  string `json:"memberLastName"`
	MemberAge       uint   `json:"memberAge"`
}

type JoinTourRequest struct {
	TourId          uint   `json:"tourId"`
	TouristUsername string `json:"touristUsername"`
	JoinedMembers   []struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Age       uint   `json:"age"`
	} `json:"joinedMembers"`
}

type JoinedMembers struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       uint   `json:"age"`
}

func GetALlJoinings() ([]Joining, error) {
	var joining []Joining
	result := db.Model(&Joining{}).Find(&joining)
	return joining, result.Error
}

func GetJoiningsByTourId(tourId uint) (joinings []Joining, err error) {
	result := db.Model(&Joining{}).Where("tour_id = ?", tourId).Find(&joinings)

	return joinings, result.Error
}

func GetJoinedMembersByTourId(tourId uint) (joinedMembers []JoinedMembers, err error) {
	joinings, err := GetJoiningsByTourId(tourId)

	if err != nil {
		return joinedMembers, err
	}

	for _, joining := range joinings {
		joinedMembers = append(joinedMembers, JoinedMembers{
			FirstName: joining.MemberFirstName,
			LastName:  joining.MemberLastName,
			Age:       joining.MemberAge,
		})
	}

	return joinedMembers, nil
}

func CreateJoining(joining *Joining) (err error) {
	result := db.Model(&Joining{}).Create(joining)

	return result.Error
}

func DeleteJoiningByTourId(tourId uint, db *gorm.DB) (err error) {
	result := db.Model(&Joining{}).Where("tour_id = ?", tourId).Delete(&Joining{})

	return result.Error
}
