package models

import "time"

type Suggestion struct {
	SuggestionId     uint      `gorm:"primaryKey;autoIncrement" json:"suggestionId"`
	Description      string    `gorm:"not null" json:"description"`
	TouristUsername  string    `gorm:"not null" json:"touristUsername"`
	SuggestTimestamp time.Time `gorm:"autoCreateTime" json:"suggestTimestamp"`
}
