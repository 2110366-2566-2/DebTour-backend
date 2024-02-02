package models

type Tourist struct {
	Username       string `json:"username"`
	CitizenId      string `json:"citizenId"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Address        string `json:"address"`
	BirthDate      string `json:"birthDate"`
	Gender         string `json:"gender"`
	DefaultPayment string `json:"defaultPayment"`
}
