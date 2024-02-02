package models

import "time"

type Suggestion struct {
	SuggestionId     uint      `json:"suggestionId"`
	Description      string    `json:"description"`
	TouristUsername  string    `json:"touristUsername"`
	SuggestTimestamp time.Time `json:"suggestTimestamp"`
}
