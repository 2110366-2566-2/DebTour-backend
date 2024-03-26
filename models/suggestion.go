package models

import "time"

type Suggestion struct {
	SuggestionId     uint      `gorm:"primaryKey;autoIncrement" json:"suggestionId"`
	Description      string    `gorm:"not null" json:"description"`
	TouristUsername  string    `gorm:"not null" json:"touristUsername"`
	SuggestTimestamp time.Time `gorm:"autoCreateTime" json:"suggestTimestamp"`
}

type SuggestionRequest struct {
	Description     string          `json:"description"`
	TouristUsername string          `json:"touristUsername"`
	LocationRequest LocationRequest `json:"locationRequest"`
}

type SuggestionWithLocation struct {
	SuggestionId     uint      `json:"suggestionId"`
	Description      string    `json:"description"`
	TouristUsername  string    `json:"touristUsername"`
	SuggestTimestamp time.Time `json:"suggestTimestamp"`
	Location         Location  `json:"location"`
}

func ToSuggestionWithLocation(suggestion Suggestion, location Location) *SuggestionWithLocation {
	return &SuggestionWithLocation{
		SuggestionId:     suggestion.SuggestionId,
		Description:      suggestion.Description,
		TouristUsername:  suggestion.TouristUsername,
		SuggestTimestamp: suggestion.SuggestTimestamp,
		Location:         location,
	}
}

func ToSuggestion(suggestionRequest SuggestionRequest) *Suggestion {
	return &Suggestion{
		Description:     suggestionRequest.Description,
		TouristUsername: suggestionRequest.TouristUsername,
	}
}
