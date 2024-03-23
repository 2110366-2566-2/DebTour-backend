package database

import (
	"DebTour/models"
	"gorm.io/gorm"
)

func GetLocationById(locationId uint, db *gorm.DB) (location models.Location, err error) {

	// find location by id in the database
	result := db.Model(&models.Location{}).Where("location_id = ?", locationId).First(&location)

	return location, result.Error
}

func CreateLocation(location *models.Location, db *gorm.DB) (err error) {

	// create location in the database
	result := db.Model(&models.Location{}).Create(location)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateLocation(location *models.Location, db *gorm.DB) (err error) {

	// check if location exists
	if _, err = GetLocationById(location.LocationId, db); err != nil {
		return err
	}

	// update location in the database
	result := db.Model(&models.Location{}).Where("location_id = ?", location.LocationId).Updates(location)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteLocation(locationId uint, db *gorm.DB) (err error) {

	// check if location exists
	if _, err = GetLocationById(locationId, db); err != nil {
		return err
	}

	// delete location in the database
	result := db.Model(&models.Location{}).Where("location_id = ?", locationId).Delete(&models.Location{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}