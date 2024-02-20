package models

type Tourist struct {
	Username       string `gorm:"primary_key" json:"username"`
	CitizenId      string `gorm:"unique" json:"citizenId"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Address        string `json:"address"`
	BirthDate      string `json:"birthDate"`
	Gender         string `json:"Gender"`
	DefaultPayment string `json:"defaultPayment"`
}
