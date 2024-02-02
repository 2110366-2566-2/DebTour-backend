package models

type Review struct {
	ReviewId    uint    `json:"reviewId"`
	Description *string `json:"description"`
	RatingScore uint    `json:"ratingScore"`
}
