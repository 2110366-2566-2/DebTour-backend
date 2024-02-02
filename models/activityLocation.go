package models

type ActivityLocation struct {
	TourId     uint `json:"tourId"`
	ActivityId uint `json:"activityId"`
	LocationId uint `json:"locationId"`
}
