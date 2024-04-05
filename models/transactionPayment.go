package models

import "time"

type TransactionPayment struct {
	TransactionId   uint      `gorm:"foreignKey;not null" json:"transactionId"`
	StripeID        string    `gorm:"not null" json:"stripeID"`
	TourId          uint      `gorm:"foreignKey;not null" json:"tourId"`
	TouristUsername string    `gorm:"foreignKey;not null" json:"touristUsername"`
	Timestamp       time.Time `gorm:"autoCreateTime" json:"timestamp"`
}

type TransactionPaymentCreateForm struct {
	TourId          uint    `json:"tourId"`
	TouristUsername string  `json:"touristUsername"`
	Amount          float64 `json:"amount"`
	Method          string  `json:"method"`
	Status          string  `json:"status"`
}

func ToTransactionPayment(transaction Transaction, form TransactionPaymentCreateForm) TransactionPayment {
	return TransactionPayment{
		TransactionId:   transaction.TransactionId,
		StripeID:        transaction.StripeID,
		TourId:          form.TourId,
		TouristUsername: form.TouristUsername,
		Timestamp:       time.Now(),
	}
}

type FullTransactionPayment struct {
	TransactionId   uint      `gorm:"foreignKey;not null" json:"transactionId"`
	StripeID        string    `gorm:"not null" json:"stripeID"`
	TourId          uint      `gorm:"foreignKey;not null" json:"tourId"`
	TouristUsername string    `gorm:"foreignKey;not null" json:"touristUsername"`
	Timestamp       time.Time `gorm:"autoCreateTime" json:"timestamp"`
	Amount          float64   `json:"amount"`
	Method          string    `json:"method"`
	Status          string    `json:"status"`
}

func ToFullTransactionPayment(transaction Transaction, transactionPayment TransactionPayment) FullTransactionPayment {
	return FullTransactionPayment{
		TransactionId:   transactionPayment.TransactionId,
		StripeID:        transaction.StripeID,
		TourId:          transactionPayment.TourId,
		TouristUsername: transactionPayment.TouristUsername,
		Timestamp:       transactionPayment.Timestamp,
		Amount:          transaction.Amount,
		Method:          transaction.Method,
		Status:          transaction.Status,
	}
}
