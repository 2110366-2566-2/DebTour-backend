package models

import "time"

type TransactionPayment struct {
	TransactionId   uint      `gorm:"foreignKey;not null" json:"transactionId"`
	TourId          uint      `gorm:"foreignKey;not null" json:"tourId"`
	TouristUsername string    `gorm:"foreignKey;not null" json:"touristUsername"`
	Timestamp       time.Time `gorm:"autoCreateTime" json:"timestamp"`
}
