package database

import (
	"DebTour/models"
	"gorm.io/gorm"
)

func GetAllReviews(db *gorm.DB) (reviews []models.Review, err error) {
	result := db.Model(&models.Review{}).Find(&reviews)

	if result.Error != nil {
		return nil, result.Error
	}

	return reviews, nil
}

func GetReviewById(reviewId uint, db *gorm.DB) (review models.Review, err error) {
	result := db.Model(&models.Review{}).First(&review, reviewId)

	if result.Error != nil {
		return models.Review{}, result.Error
	}

	return review, nil
}

func GetReviewsByTourId(tourId uint, db *gorm.DB) (reviews []models.Review, err error) {
	result := db.Model(&models.Review{}).Where("tour_id = ?", tourId).Find(&reviews)

	if result.Error != nil {
		return nil, result.Error
	}

	return reviews, nil
}

func CreateReview(review *models.Review, db *gorm.DB) (err error) {
	result := db.Model(&models.Review{}).Create(review)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateReview(review *models.Review, db *gorm.DB) (err error) {
	result := db.Model(&models.Review{}).Save(review)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteReview(reviewId uint, db *gorm.DB) (err error) {
	result := db.Model(&models.Review{}).Delete(&models.Review{}, reviewId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteReviewsByTourId(tourId uint, db *gorm.DB) (err error) {
	result := db.Model(&models.Review{}).Where("tour_id = ?", tourId).Delete(&models.Review{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}