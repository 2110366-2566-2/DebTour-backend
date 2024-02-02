package models

import "time"

type CommentReview struct {
	ReviewId        uint      `json:"reviewId"`
	TourId          uint      `json:"tourId"`
	TouristUsername string    `json:"touristUsername"`
	Timestamp       time.Time `json:"timestamp"`
}
