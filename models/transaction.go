package models

type Transaction struct {
	TransactionId uint    `gorm:"primaryKey;autoIncrement" json:"transactionId"`
	Amount        float64 `gorm:"check:amount > 0" json:"amount"`
	Status        string  `json:"status"`
	Method        string  `json:"method"`
}
