package models

type Review struct {
	ReviewId    uint    `gorm:"primaryKey;autoIncrement" json:"reviewId"`
	TourId	  uint    `gorm:"not null" json:"tourId"`
	TouristUsername string    `gorm:"not null" json:"touristUsername"`
	Description *string `json:"description"`
	RatingScore uint    `gorm:"not null;check:rating_score >= 1 AND rating_score <= 5" json:"ratingScore"`
}

type ReviewRequest struct {
	TouristUsername string    `json:"touristUsername"`
	Description *string `json:"description"`
	RatingScore uint    `json:"ratingScore"`
}

func (r *ReviewRequest) ToReview(tourId uint) Review {
	return Review{
		TourId: tourId,
		TouristUsername: r.TouristUsername,
		Description: r.Description,
		RatingScore: r.RatingScore,
	}
}