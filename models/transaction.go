package models

type Transaction struct {
	TransactionId uint    `gorm:"primaryKey;autoIncrement" json:"transactionId"`
	Amount        float64 `gorm:"not null;check:amount > 0" json:"amount"`
	Status        string  `gorm:"not null" json:"status"`
	Method        string  `gorm:"not null" json:"method"`
}
