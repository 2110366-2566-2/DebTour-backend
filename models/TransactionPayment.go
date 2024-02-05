package models

import "time"

type TransactionPayment struct {
	TransactionId   uint      `gorm:"primaryKey" json:"transactionId"`
	TourId          uint      `gorm:"primaryKey" json:"tourId"`
	TouristUsername string    `gorm:"primaryKey" json:"touristUsername"`
	Timestamp       time.Time `gorm:"autoCreateTime" json:"timestamp"`
}
