package database

import (
	"DebTour/models"
	"gorm.io/gorm"
)

func GetAllReviews(db *gorm.DB) ([]models.Review, error) {
	var reviews []models.Review
	err := db.Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func GetReviewById(reviewId uint, db *gorm.DB) (models.Review, error) {
	var review models.Review
	err := db.First(&review, reviewId).Error
	if err != nil {
		return models.Review{}, err
	}

	return review, nil
}

func GetReviewsByTourId(tourId uint, db *gorm.DB) ([]models.Review, error) {
	var reviews []models.Review
	err := db.Where("tour_id = ?", tourId).Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func GetReviewsByTouristUsername(username string, db *gorm.DB) ([]models.Review, error) {
	var reviews []models.Review
	err := db.Where("tourist_username = ?", username).Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func GetReviewByTourIdAndTouristUsername(tourId uint, username string, db *gorm.DB) (models.Review, error) {
	var review models.Review
	err := db.Where("tour_id = ? AND tourist_username = ?", tourId, username).First(&review).Error
	if err != nil {
		return models.Review{}, err
	}

	return review, nil

}

func CreateReview(review models.Review, db *gorm.DB) error {
	return db.Create(&review).Error
}

func UpdateReview(review models.Review, db *gorm.DB) error {
	return db.Save(&review).Error
}

func DeleteReview(reviewId uint, db *gorm.DB) error {
	// check if review exists
	if _, err := GetReviewById(reviewId, db) ; err != nil {
		return err
	}

	return db.Delete(&models.Review{}, reviewId).Error
}

func DeleteReviewsByTourId(tourId uint, db *gorm.DB) error {
	// check if tour exists
	if _, err := GetTourByTourId(int(tourId), db) ; err != nil {
		return err
	}

	return db.Where("tour_id = ?", tourId).Delete(&models.Review{}).Error
}

func DeleteReviewsByTouristUsername(username string, db *gorm.DB) error {
	// check if tourist exists
	//if _, err := GetUserByUsername(username, db) ; err != nil {
	//	return err
	//}

	return db.Where("tourist_username = ?", username).Delete(&models.Review{}).Error
}