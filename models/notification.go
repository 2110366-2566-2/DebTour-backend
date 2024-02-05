package models

import (
	"time"
)

type Notification struct {
	NotificationId uint      `gorm:"primaryKey" json:"notification_id"`
	Message        string    `json:"message"`
	URL            string    `json:"url"`
	Username       string    `json:"username"`
	Timestamp      time.Time `gorm:"autoCreateTime"`
}
