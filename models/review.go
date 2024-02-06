package models

type Review struct {
	ReviewId    uint    `gorm:"primaryKey;type:SERIAL" json:"reviewId"`
	Description *string `json:"description"`
	RatingScore uint    `gorm:"check:rating_score >= 1 AND rating_score <= 5" json:"ratingScore"`
}
