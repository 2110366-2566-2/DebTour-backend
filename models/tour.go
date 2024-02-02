package models

import "time"

type Tour struct {
	TourId           uint      `json:"tourId"`
	Name             string    `json:"name"`
	StartDate        string    `json:"startDate"`
	EndDate          string    `json:"endDate"`
	Description      string    `json:"description"`
	OverviewLocation string    `json:"overviewLocation"`
	Price            float64   `json:"price"`
	RefundDueDate    string    `json:"refundDueDate"`
	MaxMemberCount   uint      `json:"maxMemberCount"`
	MemberCount      uint      `json:"memberCount"`
	Status           string    `json:"status"`
	AgencyUsername   string    `json:"agencyUsername"`
	CreatedTimestamp time.Time `json:"createdTimestamp"`
}
