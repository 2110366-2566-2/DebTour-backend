package models

type Joining struct {
	TourId          uint   `json:"tourId"`
	TouristUsername string `json:"touristUsername"`
	MemberName      string `json:"memberName"`
	Age             uint   `json:"age"`
}
