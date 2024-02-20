package models

type SuggestionLocation struct {
	SuggestionId uint `gorm:"foreignKey;not null" json:"suggestionId"`
	LocationId   uint `gorm:"foreignKey;not null" json:"locationId"`
}
