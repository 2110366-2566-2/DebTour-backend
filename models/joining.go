package models

type Joining struct {
	TourId          uint   `gorm:"foreignKey" json:"tourId"`
	TouristUsername string `gorm:"foreignKey" json:"touristUsername"`
	MemberFirstName string `json:"memberFirstName"`
	MemberLastName  string `json:"memberLastName"`
	MemberAge       uint   `json:"memberAge"`
}

func CreateJoining(joining Joining) error {
	result := db.Create(&joining)
	return result.Error
}

func GetALlJoinings() ([]Joining, error) {
	var joining []Joining
	result := db.Model(&Joining{}).Find(&joining)
	return joining, result.Error
}
