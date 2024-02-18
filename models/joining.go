package models

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
