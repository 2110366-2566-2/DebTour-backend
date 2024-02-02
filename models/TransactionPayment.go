package models

import "time"

type TransactionPayment struct {
	TransactionId   uint   `json:"transactionId"`
	TourId          uint   `json:"tourId"`
	TouristUsername string `json:"touristUsername"`
	Timestamp       time.Time
}
