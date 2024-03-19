package database

import (
	"DebTour/models"
	"encoding/base64"

	"gorm.io/gorm"
)

func GetAllTours(db *gorm.DB) (tours []models.Tour, err error) {
	// Get all tours from the database
	result := db.Find(&tours)

	return tours, result.Error
}

func GetTourWithActivitiesWithLocationWithImagesByTourId(tourId int, db *gorm.DB) (tourWithActivitiesWithLocationWithImages models.TourWithActivitiesWithLocationWithImages, err error) {
	var tour models.Tour
	result := db.First(&tour, tourId)

	if result.Error != nil {
		return tourWithActivitiesWithLocationWithImages, result.Error
	}

	activitiesWithLocation, err := GetAllActivitiesWithLocationByTourId(tour.TourId, db)

	if err != nil {
		return tourWithActivitiesWithLocationWithImages, err
	}

	images, err := GetTourImages(tour.TourId, db)

	return models.ToTourWithActivitiesWithLocationWithImages(tour, activitiesWithLocation, images)
}

func GetTourByTourId(tourId int, db *gorm.DB) (tour models.Tour, err error) {
	result := db.First(&tour, tourId)

	return tour, result.Error
}

func GetToursByAgencyUsername(username string, db *gorm.DB) (tours []models.Tour, err error) {
	result := db.Model(&models.Tour{}).Where("agency_username = ?", username).Find(&tours)

	return tours, result.Error

}

func CreateTour(tour *models.Tour, activitiesWithLocationRequest []models.ActivityWithLocationRequest, images []string, db *gorm.DB) (err error) {
	tx := db.SavePoint("BeforeCreateTour")

	result := tx.Model(&models.Tour{}).Create(tour)
	if result.Error != nil {
		tx.RollbackTo("BeforeCreateTour")
		return result.Error
	}

	for _, activityWithLocationRequest := range activitiesWithLocationRequest {
		activity := models.ToActivity(activityWithLocationRequest, tour.TourId)
		location := models.ToLocation(activityWithLocationRequest.LocationRequest, 0)
		err = CreateActivity(&activity, &location, tx)

		if err != nil {
			tx.RollbackTo("BeforeCreateTour")
			return err
		}
	}

	for _, image := range images {
		image, err := base64.StdEncoding.DecodeString(image)
		if err != nil {
			tx.RollbackTo("BeforeCreateTour")
			return err
		}
		tourImage := models.TourImage{
			TourId: tour.TourId,
			Image: image,
		}
		err = CreateTourImage(&tourImage, tx)

		if err != nil {
			tx.RollbackTo("BeforeCreateTour")
			return err
		}

	}

	return nil
}

func UpdateTour(tour *models.Tour, db *gorm.DB) (err error) {
	_, err = GetTourByTourId(int(tour.TourId), db)

	if err != nil {
		return err
	}

	result := db.Model(&models.Tour{}).Where("tour_id = ?", tour.TourId).Updates(tour)

	return result.Error
}

func DeleteTour(tourId uint, db *gorm.DB) (err error) {
	tx := db.SavePoint("BeforeDeleteTour")

	// Delete all joinings of the tour by calling the function from joining.go
	err = DeleteAllJoiningsByTourId(tourId, tx)
	if err != nil {
		tx.RollbackTo("BeforeDeleteTour")
		return err
	}

	// Delete all activities of the tour by calling the function from activity.go
	err = DeleteAllActivitiesByTourId(tourId, tx)
	if err != nil {
		tx.RollbackTo("BeforeDeleteTour")
		return err
	}

	// Delete all joinings of the tour
	err = DeleteAllJoiningsByTourId(tourId, tx)
	if err != nil {
		tx.RollbackTo("BeforeDeleteTour")
		return err
	}

	// Delete all images of the tour
	err = DeleteTourImages(tourId, tx)
	if err != nil {
		tx.RollbackTo("BeforeDeleteTour")
		return err
	}

	// Delete the tour
	result := tx.Model(&models.Tour{}).Where("tour_id = ?", tourId).Delete(&models.Tour{})
	if result.Error != nil {
		tx.RollbackTo("BeforeDeleteTour")
		return result.Error
	}

	return nil
}

func FilterTours(name, startDate, endDate, overviewLocation, memberCountFrom, memberCountTo, maxMemberCountFrom, maxMemberCountTo, availableMemberCountFrom, availableMemberCountTo, priceFrom, priceTo string, offset, limit int, db *gorm.DB) ([]models.Tour, error) {
	var tours []models.Tour
	result := db.Model(&models.Tour{}).Select("tour_id, name, start_date, end_date, overview_location, member_count, max_member_count, price").Where("name LIKE ? AND start_date >= ? AND end_date <= ? AND overview_location LIKE ? AND member_count >= ? AND member_count <= ? AND max_member_count >= ? AND max_member_count <= ? AND max_member_count - member_count >= ? AND max_member_count - member_count <= ? AND price >= ? AND price <= ?", name, startDate, endDate, overviewLocation, memberCountFrom, memberCountTo, maxMemberCountFrom, maxMemberCountTo, availableMemberCountFrom, availableMemberCountTo, priceFrom, priceTo).Limit(limit).Offset(offset).Find(&tours)
	return tours, result.Error
}

func UpdateActivitiesByTourId(tourId uint, activitiesWithLocation *[]models.ActivityWithLocation, db *gorm.DB) (err error) {
	tx := db.SavePoint("BeforeUpdateActivitiesByTourId")

	for _, activityWithLocation := range *activitiesWithLocation {
		activity := models.BackToActivity(activityWithLocation)
		activity.TourId = tourId

		err = UpdateActivity(&activity, tx)

		if err != nil {
			tx.RollbackTo("BeforeUpdateActivitiesByTourId")
			return err
		}

		// if the activity has a location, update the location
		if activityWithLocation.Location != (models.Location{}) {
			err = UpdateLocation(&activityWithLocation.Location, tx)

			if err != nil {
				tx.RollbackTo("BeforeUpdateActivitiesByTourId")
				return err
			}
		}

	}

	return nil
}

func CreateTourActivities(tourId uint, activitiesWithLocationRequest []models.ActivityWithLocationRequest, db *gorm.DB) (err error) {
	tx := db.SavePoint("BeforeCreateTourActivities")

	for _, activityWithLocationRequest := range activitiesWithLocationRequest {
		activity := models.ToActivity(activityWithLocationRequest, tourId)
		location := models.ToLocation(activityWithLocationRequest.LocationRequest, 0)
		err = CreateActivity(&activity, &location, tx)

		if err != nil {
			tx.RollbackTo("BeforeCreateTourActivities")
			return err
		}
	}

	return nil
}
