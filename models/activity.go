package models

import (
	"time"
)

type Activity struct {
	TourId         uint      `gorm:"foreignKey;not null" json:"tourId"`
	ActivityId     uint      `gorm:"primaryKey;autoIncrement" json:"activityId"`
	Name           string    `gorm:"not null" json:"name"`
	Description    *string   `gorm:"not null" json:"description"`
	StartTimestamp time.Time `gorm:"not null" json:"startTimestamp"`
	EndTimestamp   time.Time `gorm:"not null" json:"endTimestamp"`
}

type ActivityWithLocationRequest struct {
	Name           string    `json:"name"`
	Description    *string   `json:"description"`
	StartTimestamp time.Time `json:"startTimestamp"`
	EndTimestamp   time.Time `json:"endTimestamp"`
	LocationRequest       LocationRequest  `json:"location"`
}

func ToActivity(request ActivityWithLocationRequest, tourId uint) Activity {
	return Activity{
		TourId:         tourId,
		Name:           request.Name,
		Description:    request.Description,
		StartTimestamp: request.StartTimestamp,
		EndTimestamp:   request.EndTimestamp,
	}
}

type ActivityWithLocation struct {
	TourId		 uint      `gorm:"foreignKey;not null" json:"tourId"`
	ActivityId     uint      `gorm:"primaryKey;autoIncrement" json:"activityId"`
	Name           string    `gorm:"not null" json:"name"`
	Description    *string   `gorm:"not null" json:"description"`
	StartTimestamp time.Time `gorm:"not null" json:"startTimestamp"`
	EndTimestamp   time.Time `gorm:"not null" json:"endTimestamp"`
	Location       Location  `json:"location"`
}

func ToActivityWithLocation(activity Activity, location Location) ActivityWithLocation {
	return ActivityWithLocation{
		TourId:         activity.TourId,
		ActivityId:     activity.ActivityId,
		Name:           activity.Name,
		Description:    activity.Description,
		StartTimestamp: activity.StartTimestamp,
		EndTimestamp:   activity.EndTimestamp,
		Location: location,
	}
}

func BackToActivity(activityWithLocation ActivityWithLocation) Activity {
	return Activity{
		TourId:         activityWithLocation.TourId,
		ActivityId:     activityWithLocation.ActivityId,
		Name:           activityWithLocation.Name,
		Description:    activityWithLocation.Description,
		StartTimestamp: activityWithLocation.StartTimestamp,
		EndTimestamp:   activityWithLocation.EndTimestamp,
	}
}