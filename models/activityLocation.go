package models

type ActivityLocation struct {
	TourId     uint `gorm:"foreignKey;not null" json:"tourId"`
	ActivityId uint `gorm:"foreignKey;not null" json:"activityId"`
	LocationId uint `gorm:"foreignKey;not null" json:"locationId"`
}

func CreateActivityLocation(activityLocation *ActivityLocation) (err error) {
	result := db.Model(&ActivityLocation{}).Create(activityLocation)

	return result.Error
}

func GetActivityLocationByTourId(tourId uint) (activityLocations []ActivityLocation, err error) {
	result := db.Model(&ActivityLocation{}).Where("tour_id = ?", tourId).Find(&activityLocations)

	return activityLocations, result.Error
}

func GetActivityLocationByActivityId(activityId uint) (activityLocations ActivityLocation, err error) {
	result := db.Model(&ActivityLocation{}).Where("activity_id = ?", activityId).First(&activityLocations)

	return activityLocations, result.Error
}
