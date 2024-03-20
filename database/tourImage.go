package database

import (
	"DebTour/models"
	"encoding/base64"
	"gorm.io/gorm"
)

func GetTourImages(tourId uint, db *gorm.DB) ([]string, error) {
	var images []models.TourImage
	err := db.Model(&models.TourImage{}).Where("tour_id = ?", tourId).Find(&images).Error
	if err != nil {
		return nil, err
	}

	var imageStrings []string
	for _, image := range images {
		encodedImage := base64.StdEncoding.EncodeToString(image.Image)
		imageStrings = append(imageStrings, encodedImage)
	}
	return imageStrings, nil
}

func CreateTourImage(tourImage *models.TourImage, db *gorm.DB) error {
	db.SavePoint("BeforeCreateTourImage")

	err := db.Create(&tourImage).Error
	if err != nil {
		db.RollbackTo("BeforeCreateTourImage")
		return err
	}

	return nil
}

func DeleteTourImage(tourId uint, image string, db *gorm.DB) error {
	err := db.Where("tour_id = ? AND image = ?", tourId, image).Delete(&models.TourImage{}).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteTourImages(tourId uint, db *gorm.DB) error {
	err := db.Where("tour_id = ?", tourId).Delete(&models.TourImage{}).Error
	if err != nil {
		return err
	}
	return nil
}
