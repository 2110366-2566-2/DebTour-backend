package models

type Joining struct {
	TourId          uint   `gorm:"primaryKey" json:"tourId"`
	TouristUsername string `gorm:"primaryKey" json:"touristUsername"`
	MemberName      string `gorm:"primaryKey" json:"memberName"`
	Age             uint   `json:"age"`
}
