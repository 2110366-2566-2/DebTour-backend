package database

import (
	"DebTour/models"
	"gorm.io/gorm"
)
//Checked
func GetAllActivities(db *gorm.DB) (activities []models.Activity, err error) {
	// find all activities in the database
	result := db.Model(&models.Activity{}).Find(&activities)

	return activities, result.Error
}
//Checked
func GetActivityById(activityId uint, db *gorm.DB) (activity models.Activity, err error) {
	// find activity by id in the database
	result := db.Model(&models.Activity{}).First(&activity, activityId)

	return activity, result.Error
}

//Checked
func GetActivityWithLocationById(activityId uint, db *gorm.DB) (activityWithLocation models.ActivityWithLocation, err error) {
	var activity models.Activity
	// find activity by id in the database
	result := db.Model(&models.Activity{}).First(&activity, activityId)

	if result.Error != nil {
		return models.ActivityWithLocation{}, result.Error
	}

	var activityLocation models.ActivityLocation
	// find activity location by activity id in the database
	activityLocation, err = GetActivityLocationByActivityId(activityId, db)

	if err != nil {
		return models.ActivityWithLocation{}, err
	}

	var location models.Location
	// find location by location id in the database
	location, err = GetLocationById(activityLocation.LocationId, db)

	if err != nil {
		return models.ActivityWithLocation{}, err
	}

	activityWithLocation = models.ToActivityWithLocation(activity, location)

	return activityWithLocation, result.Error
}
//Checked
func GetAllActivitiesWithLocationByTourId(tourId uint, db *gorm.DB) (activitiesWithLocation []models.ActivityWithLocation, err error) {

	// find all activities by tour id in the database
	var activities []models.Activity
	result := db.Model(&models.Activity{}).Where("tour_id = ?", tourId).Find(&activities)

	if result.Error != nil {
		return activitiesWithLocation, result.Error
	}

	// get all activities with location
	for _, activity := range activities {
		activityWithLocation, err := GetActivityWithLocationById(activity.ActivityId, db)

		if err != nil {
			return activitiesWithLocation, err
		}

		activitiesWithLocation = append(activitiesWithLocation, activityWithLocation)
	}

	return activitiesWithLocation, result.Error
}
//Checked
func CreateActivity(activity *models.Activity, location *models.Location, db *gorm.DB) (err error) {
	tx := db.SavePoint("BeforeCreateActivity")

	// create activity in the database
	result := tx.Model(&models.Activity{}).Create(activity)

	if result.Error != nil {
		tx.RollbackTo("BeforeCreateActivity")
		return result.Error
	}

	// create location in the database
	if err := CreateLocation(location, tx); err != nil {
		tx.RollbackTo("BeforeCreateActivity")
		return err
	}

	// create activity location in the database
	err = CreateActivityLocation(&models.ActivityLocation{
		TourId:     activity.TourId,
		ActivityId: activity.ActivityId,
		LocationId: location.LocationId,
	}, tx)

	if err != nil {
		tx.RollbackTo("BeforeCreateActivity")
		return err
	}

	return nil
}
//Checked
func UpdateActivity(activity *models.Activity, db *gorm.DB) (err error) {

	// check if activity exists
	if _, err = GetActivityById(activity.ActivityId, db); err != nil {
		return err
	}

	// update activity in the database
	result := db.Model(&models.Activity{}).Where("activity_id = ?", activity.ActivityId).Updates(activity)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
//Checked
func DeleteActivity(activityId uint, db *gorm.DB) (err error) {
	tx := db.SavePoint("BeforeDeleteActivity")

	// check if activity exists
	if _, err = GetActivityById(activityId, tx); err != nil {
		tx.RollbackTo("BeforeDeleteActivity")
		return err
	}

	activityWithLocation, err := GetActivityWithLocationById(activityId, tx)

	if err != nil {
		tx.RollbackTo("BeforeDeleteActivity")
		return err
	}

	// delete activity location in the database (also cascading delete with location)
	if err := DeleteActivityLocation(models.ActivityLocation{
		TourId:     activityWithLocation.TourId,
		ActivityId: activityWithLocation.ActivityId,
		LocationId: activityWithLocation.Location.LocationId,
	}, tx); err != nil {
		tx.RollbackTo("BeforeDeleteActivity")
		return err
	}

	return nil
}
//Checked
func DeleteAllActivitiesByTourId(tourId uint, db *gorm.DB) (err error) {
	tx := db.SavePoint("BeforeDeleteAllActivitiesByTourId")

	// delete all activityLocations by tour id in the database
	err = DeleteAllActivityLocationsByTourId(tourId, tx)

	if err != nil {
		tx.RollbackTo("BeforeDeleteAllActivitiesByTourId")
		return err
	}

	result := tx.Model(&models.Activity{}).Where("tour_id = ?", tourId).Delete(&models.Activity{})

	if result.Error != nil {
		tx.RollbackTo("BeforeDeleteAllActivitiesByTourId")
		return result.Error
	}

	return nil
}
