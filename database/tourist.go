package database

import (
	"DebTour/models"

	"gorm.io/gorm"
)

func GetTouristByUsername(username string, db *gorm.DB) (touristWithUser models.TouristWithUser, err error) {
	var tourist models.Tourist
	result := db.Model(&models.Tourist{}).Where("username = ?", username).First(&tourist)

	user, err := GetUserByUsername(username, db)

	if err != nil {
		return touristWithUser, err
	}

	touristWithUser = models.ToTouristWithUser(tourist, user)

	return touristWithUser, result.Error
}

func CreateTourist(tourist *models.Tourist, db *gorm.DB) (err error) {
	db.SavePoint("BeforeCreateTourist")
	result := db.Model(&models.Tourist{}).Create(tourist)
	if result.Error != nil {
		db.RollbackTo("BeforeCreateTourist")
		return result.Error
	}
	return result.Error
}

func GetAllTourists(db *gorm.DB) (touristsWithUser []models.TouristWithUser, err error) {
	var tourists []models.Tourist
	result := db.Model(&models.Tourist{}).Find(&tourists)

	for _, tourist := range tourists {
		user, err := GetUserByUsername(tourist.Username, db)
		if err != nil {
			return touristsWithUser, err
		}
		touristsWithUser = append(touristsWithUser, models.ToTouristWithUser(tourist, user))
	}

	return touristsWithUser, result.Error
}

// func delete tourist by username and db
func DeleteTouristByUsername(username string, db *gorm.DB) (err error) {
	result := db.Model(&models.Tourist{}).Where("username = ?", username).Delete(&models.Tourist{})
	return result.Error
}

func UpdateTouristByUsername(username string, tourist models.Tourist, db *gorm.DB) (err error) {
	existingUser, err := GetTouristByUsername(username, db)
	if err != nil {
		return err
	}
	existingTourist := models.ToTourist(existingUser)
	//update tourist
	result := db.Model(&existingTourist).Updates(tourist)
	return result.Error
}
