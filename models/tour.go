package models

import "time"

type Tour struct {
	TourId           uint      `gorm:"primaryKey;autoIncrement" json:"tourId"`
	Name             string    `json:"name"`
	StartDate        string    `json:"startDate"`
	EndDate          string    `json:"endDate"`
	Description      string    `json:"description"`
	OverviewLocation string    `json:"overviewLocation"`
	Price            float64   `gorm:"check:price > 0" json:"price"`
	RefundDueDate    string    `json:"refundDueDate"`
	MaxMemberCount   uint      `json:"maxMemberCount"`
	MemberCount      uint      `json:"memberCount"`
	Status           string    `json:"status"`
	AgencyUsername   string    `json:"agencyUsername"`
	CreatedTimestamp time.Time `gorm:"autoCreateTime" json:"createdTimestamp"`
}

func GetAllTours() (tours []Tour, err error) {
	result := db.Find(&tours)

	return tours, result.Error
}

func CreateTour(tour *Tour) (err error) {
	result := db.Model(&Tour{}).Create(tour)

	return result.Error
}

func UpdateTour(tour *Tour) (err error) {
	result := db.Model(&Tour{}).Where("tour_id = ?", tour.TourId).Updates(tour)

	return result.Error
}

func DeleteTour(tourId uint) (err error) {

	result := db.Model(&Tour{}).Where("tour_id = ?", tourId).Delete(&Tour{})

	return result.Error
}
