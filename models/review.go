package models

type Review struct {
	ReviewId    uint    `gorm:"primaryKey;autoIncrement" json:"reviewId"`
	TourId      uint    `gorm:"foreignKey;not null" json:"tourId"`
	Description *string `json:"description"`
	RatingScore uint    `gorm:"not null;check:rating_score >= 1 AND rating_score <= 5" json:"ratingScore"`
}

type ReviewRequest struct {
	Description *string `json:"description"`
	RatingScore uint    `json:"ratingScore"`
}

func (r *ReviewRequest) ToReview(tourId uint) *Review {
	return &Review{
		TourId:      tourId,
		Description: r.Description,
		RatingScore: r.RatingScore,
	}
}