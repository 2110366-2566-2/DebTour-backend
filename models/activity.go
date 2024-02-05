package models

import "time"

type Activity struct {
	TourId         uint      `gorm:"primaryKey" json:"tourId"`
	ActivityId     uint      `gorm:"type:SERIAL" json:"activityId"`
	Name           string    `json:"name"`
	Description    *string   `json:"description"`
	StartTimestamp time.Time `json:"startTimestamp"`
	EndTimestamp   time.Time `json:"endTimestamp"`
}
