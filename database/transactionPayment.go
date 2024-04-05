package database

import (
	"DebTour/models"
	"strconv"

	"gorm.io/gorm"
)

func GetAllTransactionPayments(db *gorm.DB) (fullTransactionPayment []models.FullTransactionPayment, err error) {
	var transactionPayments []models.TransactionPayment
	result := db.Model(&models.TransactionPayment{}).Find(&transactionPayments)

	for _, transactionPayment := range transactionPayments {
		transaction := models.Transaction{}
		err := GetTransactionByTransactionId(&transaction, strconv.Itoa(int(transactionPayment.TransactionId)), db)
		if err != nil {
			return []models.FullTransactionPayment{}, err
		}
		fullTransactionPayment = append(fullTransactionPayment, models.ToFullTransactionPayment(transaction, transactionPayment))
	}
	return fullTransactionPayment, result.Error
}

func GetTransactionPaymentByTourId(tourId string, db *gorm.DB) (fullTransactionPayment []models.FullTransactionPayment, err error) {
	var transactionPayments []models.TransactionPayment
	result := db.Model(&models.TransactionPayment{}).Where("tour_id = ?", tourId).Find(&transactionPayments)

	for _, transactionPayment := range transactionPayments {
		transaction := models.Transaction{}
		err := GetTransactionByTransactionId(&transaction, strconv.Itoa(int(transactionPayment.TransactionId)), db)
		if err != nil {
			return []models.FullTransactionPayment{}, err
		}
		fullTransactionPayment = append(fullTransactionPayment, models.ToFullTransactionPayment(transaction, transactionPayment))
	}
	return fullTransactionPayment, result.Error
}

func GetTransactionPaymentByTouristUsername(username string, db *gorm.DB) (fullTransactionPayment []models.FullTransactionPayment, err error) {
	var transactionPayments []models.TransactionPayment
	result := db.Model(&models.TransactionPayment{}).Where("tourist_username = ?", username).Find(&transactionPayments)

	for _, transactionPayment := range transactionPayments {
		transaction := models.Transaction{}
		err := GetTransactionByTransactionId(&transaction, strconv.Itoa(int(transactionPayment.TransactionId)), db)
		if err != nil {
			return []models.FullTransactionPayment{}, err
		}
		fullTransactionPayment = append(fullTransactionPayment, models.ToFullTransactionPayment(transaction, transactionPayment))
	}
	return fullTransactionPayment, result.Error

}

func GetTransactionPaymentByTransactionId(transactionId string, db *gorm.DB) (fullTransactionPayment models.FullTransactionPayment, err error) {
	var transactionPayment models.TransactionPayment
	result := db.Model(&models.TransactionPayment{}).Where("transaction_id = ?", transactionId).First(&transactionPayment)

	transaction := models.Transaction{}
	err = GetTransactionByTransactionId(&transaction, transactionId, db)
	if err != nil {
		return models.FullTransactionPayment{}, err
	}
	fullTransactionPayment = models.ToFullTransactionPayment(transaction, transactionPayment)
	return fullTransactionPayment, result.Error
}

func CreateTransactionPayment(transactionPaymentCreateForm models.TransactionPaymentCreateForm, StripeID string, db *gorm.DB) (err error) {
	db.SavePoint("BeforeCreateTransactionPayment")
	transactionCreateForm := models.ToTransactionCreateForm(transactionPaymentCreateForm, StripeID)
	transaction := models.ToTransaction(transactionCreateForm)
	err = CreateTransaction(&transaction, db)
	if err != nil {
		db.RollbackTo("BeforeCreateTransactionPayment")
		return err
	}

	transactionPayment := models.ToTransactionPayment(transaction, transactionPaymentCreateForm)
	result := db.Model(&models.TransactionPayment{}).Create(transactionPayment)

	return result.Error
}

func DeleteTransactionPayment(transactionId string, db *gorm.DB) (err error) {
	// Check if transaction exists
	if _, err := GetTransactionPaymentByTransactionId(transactionId, db); err != nil {
		return err
	}

	db.SavePoint("BeforeDeleteTransactionPayment")

	// Delete transaction
	if err := DeleteTransactionByTransactionId(transactionId, db); err != nil {
		db.RollbackTo("BeforeDeleteTransactionPayment")
		return err
	}

	result := db.Model(&models.TransactionPayment{}).Where("transaction_id = ?", transactionId).Delete(&models.TransactionPayment{})
	return result.Error
}

func DeleteTransactionPaymentByTourId(tourId string, db *gorm.DB) (err error) {
	// Check if transaction exists
	if _, err := GetTransactionPaymentByTourId(tourId, db); err != nil {
		return err
	}

	db.SavePoint("BeforeDeleteTransactionPayment")

	// Delete transaction
	if err := DeleteTransactionByTourId(tourId, db); err != nil {
		db.RollbackTo("BeforeDeleteTransactionPayment")
		return err
	}

	result := db.Model(&models.TransactionPayment{}).Where("tour_id = ?", tourId).Delete(&models.TransactionPayment{})
	return result.Error
}

func DeleteTransactionPaymentByTouristUsername(username string, db *gorm.DB) (err error) {
	// Check if transaction exists
	if _, err := GetTransactionPaymentByTouristUsername(username, db); err != nil {
		return err
	}

	db.SavePoint("BeforeDeleteTransactionPayment")

	// Delete transaction
	if err := DeleteTransactionByTouristUsername(username, db); err != nil {
		db.RollbackTo("BeforeDeleteTransactionPayment")
		return err
	}

	result := db.Model(&models.TransactionPayment{}).Where("tourist_username = ?", username).Delete(&models.TransactionPayment{})
	return result.Error
}
