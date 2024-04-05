package database

import (
	"DebTour/models"
	"gorm.io/gorm"
)

func GetTransactionByTransactionId(transaction *models.Transaction, transactionId string, db *gorm.DB) (err error) {
	result := db.Model(&models.Transaction{}).Where("transaction_id = ?", transactionId).First(transaction)
	return result.Error
}

func GetTransactionByTourId(transactions *[]models.Transaction, tourId string, db *gorm.DB) (err error) {
	result := db.Model(&models.Transaction{}).Where("tour_id = ?", tourId).Find(transactions)
	return result.Error
}

func GetTransactionByTouristUsername(transactions *[]models.Transaction, username string, db *gorm.DB) (err error) {
	result := db.Model(&models.Transaction{}).Where("tourist_username = ?", username).Find(transactions)
	return result.Error
}

func CreateTransaction(transaction *models.Transaction, db *gorm.DB) (err error) {
	result := db.Model(&models.Transaction{}).Create(transaction)
	return result.Error
}

func DeleteTransactionByTransactionId(transactionId string, db *gorm.DB) (err error) {
	if err := GetTransactionByTransactionId(&models.Transaction{}, transactionId, db); err != nil {
		return err
	}

	result := db.Model(&models.Transaction{}).Where("transaction_id = ?", transactionId).Delete(&models.Transaction{})
	return result.Error
}

func DeleteTransactionByTourId(tourId string, db *gorm.DB) (err error) {
	if err := GetTransactionByTourId(&[]models.Transaction{}, tourId, db); err != nil {
		return err
	}

	result := db.Model(&models.Transaction{}).Where("tour_id = ?", tourId).Delete(&models.Transaction{})
	return result.Error
}

func DeleteTransactionByTouristUsername(username string, db *gorm.DB) (err error) {
	if err := GetTransactionByTouristUsername(&[]models.Transaction{}, username, db); err != nil {
		return err
	}

	result := db.Model(&models.Transaction{}).Where("tourist_username = ?", username).Delete(&models.Transaction{})
	return result.Error
}

func UpdateTransactionStatus(transactionId string, status string, db *gorm.DB) (err error) {
	transaction := models.Transaction{Status: status}
	result := db.Model(&models.Transaction{}).Where("transaction_id = ?", transactionId).Updates(transaction)
	return result.Error
}