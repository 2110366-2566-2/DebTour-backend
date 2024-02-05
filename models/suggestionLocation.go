package models

type SuggestionLocation struct {
	SuggestionId uint `gorm:"primaryKey" json:"suggestionId"`
	LocationId   uint `gorm:"primaryKey" json:"locationId"`
}
