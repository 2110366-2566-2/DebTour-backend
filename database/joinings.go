package database

import (
	"DebTour/models"
	"gorm.io/gorm"
)

func GetALlJoinings(db *gorm.DB) (joinings []models.Joining, err error) {
	//	Get all joinings from the database
	result := db.Find(&joinings)

	return joinings, result.Error
}

func GetJoiningsByTourId(tourId uint, db *gorm.DB) (joinings []models.Joining, err error) {
	result := db.Model(&models.Joining{}).Where("tour_id = ?", tourId).Find(&joinings)

	return joinings, result.Error
}

func GetJoinedMembersByTourId(tourId uint, db *gorm.DB) (joinedMembers []models.JoinedMembers, err error) {
	joinings, err := GetJoiningsByTourId(tourId, db)

	if err != nil {
		return joinedMembers, err
	}

	for _, joining := range joinings {
		joinedMembers = append(joinedMembers, models.JoinedMembers{
			FirstName: joining.MemberFirstName,
			LastName:  joining.MemberLastName,
			Age:       joining.MemberAge,
		})
	}

	return joinedMembers, nil
}

func CreateJoining(joining *models.Joining, db *gorm.DB) (err error) {
	result := db.Model(&models.Joining{}).Create(joining)

	return result.Error
}

func DeleteAllJoiningsByTourId(tourId uint, db *gorm.DB) (err error) {
	result := db.Model(&models.Joining{}).Where("tour_id = ?", tourId).Delete(&models.Joining{})

	return result.Error
}
