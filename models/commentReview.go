package models

import "time"

type CommentReview struct {
	ReviewId        uint      `gorm:"foreignKey;not null" json:"reviewId"`
	TourId          uint      `gorm:"foreignKey;not null" json:"tourId"`
	TouristUsername string    `gorm:"foreignKey;not null" json:"touristUsername"`
	Timestamp       time.Time `gorm:"autoCreateTime" json:"timestamp"`
}
