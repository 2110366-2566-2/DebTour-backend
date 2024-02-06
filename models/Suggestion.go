package models

import "time"

type Suggestion struct {
	SuggestionId     uint      `gorm:"type:SERIAL" json:"suggestionId"`
	Description      string    `json:"description"`
	TouristUsername  string    `json:"touristUsername"`
	SuggestTimestamp time.Time `gorm:"autoCreateTime" json:"suggestTimestamp"`
}
