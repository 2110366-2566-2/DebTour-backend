package database

import (
	"DebTour/models"
	"gorm.io/gorm"
)

func GetAllActivityLocations(db *gorm.DB) (activityLocations []models.ActivityLocation, err error) {

	// find all activityLocations in the database
	result := db.Model(&models.ActivityLocation{}).Find(&activityLocations)

	return activityLocations, result.Error
}

func GetActivityLocation(tourId uint, activityId uint, locationId uint, db *gorm.DB) (activityLocation models.ActivityLocation, err error) {

	// find the activityLocation in the database
	result := db.Model(&models.ActivityLocation{}).Where("tour_id = ? AND activity_id = ? AND location_id = ?", tourId, activityId, locationId).First(&activityLocation)

	return activityLocation, result.Error
}

func GetAllActivityLocationsByTourId(tourId uint, db *gorm.DB) (activityLocations []models.ActivityLocation, err error) {

	// find all activityLocations in the database
	result := db.Model(&models.ActivityLocation{}).Where("tour_id = ?", tourId).Find(&activityLocations)

	return activityLocations, result.Error
}

func GetActivityLocationByActivityId(activityId uint, db *gorm.DB) (activityLocations models.ActivityLocation, err error) {

	// find the activityLocation in the database
	result := db.Model(&models.ActivityLocation{}).Where("activity_id = ?", activityId).First(&activityLocations)

	return activityLocations, result.Error
}

func CreateActivityLocation(activityLocation *models.ActivityLocation, db *gorm.DB) (err error) {

	// create the activityLocation in the database
	result := db.Model(&models.ActivityLocation{}).Create(activityLocation)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateActivityLocation(activityLocation *models.ActivityLocation, db *gorm.DB) (err error) {

	// check if activityLocation exists
	if _, err = GetActivityLocation(activityLocation.TourId, activityLocation.ActivityId, activityLocation.LocationId, db); err != nil {
		return err
	}

	// update the activityLocation in the database
	result := db.Model(&models.ActivityLocation{}).Where("tour_id = ? AND activity_id = ? AND location_id = ?", activityLocation.TourId, activityLocation.ActivityId, activityLocation.LocationId).Updates(activityLocation)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteActivityLocation(activityLocation models.ActivityLocation, db *gorm.DB) (err error) {
	tx := db.SavePoint("BeforeDeleteActivityLocation")

	// check if activityLocation exists
	if _, err = GetActivityLocation(activityLocation.TourId, activityLocation.ActivityId, activityLocation.LocationId, tx); err != nil {
		tx.RollbackTo("BeforeDeleteActivityLocation")
		return err
	}

	// delete the location in the database
	if err := DeleteLocation(activityLocation.LocationId, tx); err != nil {
		tx.RollbackTo("BeforeDeleteActivityLocation")
		return err
	}

	// delete the activityLocation in the database
	result := tx.Model(&models.ActivityLocation{}).Where("tour_id = ? AND activity_id = ? AND location_id = ?", activityLocation.TourId, activityLocation.ActivityId, activityLocation.LocationId).Delete(&activityLocation)

	if result.Error != nil {
		tx.RollbackTo("BeforeDeleteActivityLocation")
		return result.Error
	}

	return nil
}

func DeleteAllActivityLocationsByTourId(tourId uint, db *gorm.DB) (err error) {
	tx := db.SavePoint("BeforeDeleteAllActivityLocationsByTourId")

	activityLocations, err := GetAllActivityLocationsByTourId(tourId, tx)

	if err != nil {
		tx.RollbackTo("BeforeDeleteAllActivityLocationsByTourId")
		return err
	}

	for _, activityLocation := range activityLocations {
		// delete the location in the database
		if err := DeleteLocation(activityLocation.LocationId, tx); err != nil {
			tx.RollbackTo("BeforeDeleteAllActivityLocationsByTourId")
			return err
		}
	}

	// delete all activityLocations in the database
	result := tx.Model(&models.ActivityLocation{}).Where("tour_id = ?", tourId).Delete(&models.ActivityLocation{})

	if result.Error != nil {
		tx.RollbackTo("BeforeDeleteAllActivityLocationsByTourId")
		return result.Error
	}

	return nil
}