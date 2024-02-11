package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	TourId         uint      `gorm:"foreignKey;not null" json:"tourId"`
	ActivityId     uint      `gorm:"primaryKey;autoIncrement" json:"activityId"`
	Name           string    `gorm:"not null" json:"name"`
	Description    *string   `gorm:"not null" json:"description"`
	StartTimestamp time.Time `gorm:"not null" json:"startTimestamp"`
	EndTimestamp   time.Time `gorm:"not null" json:"endTimestamp"`
}

type ActivityRequest struct {
	Name           string          `json:"name"`
	Description    *string         `json:"description"`
	StartTimestamp time.Time       `json:"startTimestamp"`
	EndTimestamp   time.Time       `json:"endTimestamp"`
	Location       LocationRequest `json:"location"`
}

type ActivityResponse struct {
	ActivityId     uint      `json:"activityId"`
	Name           string    `json:"name"`
	Description    *string   `json:"description"`
	StartTimestamp time.Time `json:"startTimestamp"`
	EndTimestamp   time.Time `json:"endTimestamp"`
	Location       Location  `json:"location"`
}

func ToActivity(activityRequest ActivityRequest, tourId uint) Activity {
	return Activity{
		TourId:         tourId,
		Name:           activityRequest.Name,
		Description:    activityRequest.Description,
		StartTimestamp: activityRequest.StartTimestamp,
		EndTimestamp:   activityRequest.EndTimestamp,
	}
}

func ToActivityResponse(activity Activity) (ActivityResponse, error) {
	activityLocation, err := GetActivityLocationByActivityId(activity.ActivityId)

	if err != nil {
		return ActivityResponse{}, err
	}

	location, err := GetLocationById(activityLocation.LocationId)

	if err != nil {
		return ActivityResponse{}, err
	}

	return ActivityResponse{
		ActivityId:     activity.ActivityId,
		Name:           activity.Name,
		Description:    activity.Description,
		StartTimestamp: activity.StartTimestamp,
		EndTimestamp:   activity.EndTimestamp,
		Location:       location,
	}, nil
}

func GetAllActivities() (activities []Activity, err error) {
	result := db.Model(&Activity{}).Find(&activities)

	return activities, result.Error
}

func GetAllActivitiesByTourId(tourId uint) (activities []Activity, err error) {
	result := db.Model(&Activity{}).Where("tour_id = ?", tourId).Find(&activities)

	return activities, result.Error
}

func CreateActivity(activity *Activity, locationRequest LocationRequest, db *gorm.DB) (err error) {
	result := db.Model(&Activity{}).Create(activity)
	if result.Error != nil {
		return result.Error
	}

	location := ToLocation(locationRequest)
	err = CreateLocation(&location)

	if err != nil {
		return err
	}

	err = CreateActivityLocation(&ActivityLocation{
		TourId:     activity.TourId,
		ActivityId: activity.ActivityId,
		LocationId: location.LocationId,
	})

	if err != nil {
		return err
	}

	return nil
}

func GetActivityById(activityId uint) (Activity, error) {
	var activity Activity
	result := db.Model(&Activity{}).First(&activity, activityId)
	return activity, result.Error
}

func UpdateActivity(activity *Activity) (err error) {
	_, err = GetActivityById(activity.ActivityId)

	if err != nil {
		return err
	}

	result := db.Model(&Activity{}).Where("activity_id = ?", activity.ActivityId).Updates(activity)

	return result.Error
}

func DeleteActivity(activityId uint) (err error) {
	result := db.Model(&Activity{}).Where("activity_id = ?", activityId).Delete(&Activity{})

	return result.Error
}

func DeleteActivityByTourId(tourId uint) (err error) {
	result := db.Model(&Activity{}).Where("tour_id = ?", tourId).Delete(&Activity{})

	return result.Error
}
