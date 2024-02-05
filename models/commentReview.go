package models

import "time"

type CommentReview struct {
	ReviewId        uint      `gorm:"primaryKey" json:"reviewId"`
	TourId          uint      `gorm:"primaryKey" json:"tourId"`
	TouristUsername string    `gorm:"primaryKey" json:"touristUsername"`
	Timestamp       time.Time `gorm:"autoCreateTime" json:"timestamp"`
}
