package models

import "time"

type Activity struct {
	TourId         uint      `json:"tourId"`
	ActivityId     uint      `json:"activityId"`
	Name           string    `json:"name"`
	Description    *string   `json:"description"`
	StartTimestamp time.Time `json:"startTimestamp"`
	EndTimestamp   time.Time `json:"endTimestamp"`
}
