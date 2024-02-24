package models

import (
	"time"
)

type Notification struct {
	NotificationId uint      `gorm:"primaryKey" json:"notification_id"`
	Message        string    `gorm:"not null" json:"message"`
	URL            string    `gorm:"not null" json:"url"`
	Username       string    `gorm:"not null" json:"username"`
	Timestamp      time.Time `gorm:"not null" gorm:"autoCreateTime"`
}
