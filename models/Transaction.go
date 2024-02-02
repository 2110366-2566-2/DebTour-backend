package models

type Transaction struct {
	TransactionId uint    `json:"transactionId"`
	Amount        float64 `json:"amount"`
	Status        string  `json:"status"`
	Method        string  `json:"method"`
}
