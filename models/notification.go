package models

import (
	"time"
)

type Notification struct {
	NotificationId uint   `json:"notificationId"`
	Message        string `json:"message"`
	URL            string `json:"url"`
	Username       string `json:"username"`
	Timestamp      time.Time
}
