package models

type Transaction struct {
	TransactionId uint    `gorm:"primaryKey;autoIncrement" json:"transactionId"`
	StripeID      string  `gorm:"not null" json:"stripeID"`
	Amount        float64 `gorm:"not null;check:amount > 0" json:"amount"`
	Status        string  `gorm:"not null" json:"status"`
	Method        string  `gorm:"not null" json:"method"`
}

type TransactionCreateForm struct {
	StripeID string  `json:"stripeID" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
	Status   string  `json:"status" binding:"required"`
	Method   string  `json:"method" binding:"required"`
}

func ToTransaction(form TransactionCreateForm) Transaction {
	return Transaction{
		StripeID: form.StripeID,
		Amount:   form.Amount,
		Status:   form.Status,
		Method:   form.Method,
	}
}

func ToTransactionCreateForm(form TransactionPaymentCreateForm, StripeID string) TransactionCreateForm {
	return TransactionCreateForm{
		StripeID: StripeID,
		Amount:   form.Amount,
		Status:   form.Status,
		Method:   form.Method,
	}
}
