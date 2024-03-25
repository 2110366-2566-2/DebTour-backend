package database

import (
	"DebTour/models"

	"gorm.io/gorm"
)

func CreateSuggestion(suggestion *models.Suggestion, db *gorm.DB) (err error) {
	db.SavePoint("BeforeCreateSuggestion")
	err = db.Create(suggestion).Error
	if err != nil {
		db.RollbackTo("BeforeCreateSuggestion")
		return err
	}
	return nil
}

// GetSuggestionBySuggestionId
func GetSuggestionBySuggestionId(suggestionId uint, db *gorm.DB) (suggestion *models.Suggestion, err error) {
	err = db.Model(&models.Suggestion{}).Where("suggestion_id = ?", suggestionId).First(&suggestion).Error
	if err != nil {
		return nil, err
	}
	return suggestion, nil
}

func GetAllSuggestions(db *gorm.DB) (suggestions []models.Suggestion, err error) {
	err = db.Find(&suggestions).Error
	if err != nil {
		return nil, err
	}
	return suggestions, nil
}

func UpdateSuggestion(suggestionId uint, suggestion *models.Suggestion, db *gorm.DB) (err error) {
	db.SavePoint("BeforeUpdateSuggestion")
	err = db.Model(&models.Suggestion{}).Where("suggestion_id = ?", suggestionId).Updates(suggestion).Error
	if err != nil {
		db.RollbackTo("BeforeUpdateSuggestion")
		return err
	}
	return nil
}

func DeleteSuggestion(suggestionId uint, db *gorm.DB) (err error) {
	db.SavePoint("BeforeDeleteSuggestion")
	err = db.Where("suggestion_id = ?", suggestionId).Delete(&models.Suggestion{}).Error
	if err != nil {
		db.RollbackTo("BeforeDeleteSuggestion")
		return err
	}
	return nil
}

// delete suggestion by tourist username
func DeleteSuggestionByTouristUsername(touristUsername string, db *gorm.DB) (err error) {
	db.SavePoint("BeforeDeleteSuggestionByTouristUsername")
	err = db.Where("tourist_username = ?", touristUsername).Delete(&models.Suggestion{}).Error
	if err != nil {
		db.RollbackTo("BeforeDeleteSuggestionByTouristUsername")
		return err
	}
	return nil
}

func GetSuggestionWithLocationBySuggestionId(suggestionId uint, db *gorm.DB) (suggestionWithLocation *models.SuggestionWithLocation, err error) {
	var suggestion models.Suggestion
	// find suggestion by id in the database
	result := db.Model(&models.Suggestion{}).First(&suggestion, suggestionId)

	if result.Error != nil {
		return nil, result.Error
	}

	var suggestionLocation *models.SuggestionLocation
	// find suggestion location by suggestion id in the database
	suggestionLocation, err = GetSuggestionLocationBySuggestionId(suggestionId, db)

	if err != nil {
		return nil, err
	}

	var location models.Location
	// find location by location id in the database
	location, err = GetLocationById(suggestionLocation.LocationId, db)

	if err != nil {
		return nil, err
	}

	suggestionWithLocation = models.ToSuggestionWithLocation(suggestion, location)

	return suggestionWithLocation, result.Error
}

// get all suggestions with location
func GetAllSuggestionsWithLocation(db *gorm.DB) (suggestionsWithLocation []models.SuggestionWithLocation, err error) {
	// find all suggestions in the database
	var suggestions []models.Suggestion
	suggestions, err = GetAllSuggestions(db)

	if err != nil {
		return nil, err
	}

	// get all suggestions with location
	for _, suggestion := range suggestions {
		suggestionWithLocation, err := GetSuggestionWithLocationBySuggestionId(suggestion.SuggestionId, db)
		if err != nil {
			return nil, err
		}
		suggestionsWithLocation = append(suggestionsWithLocation, *suggestionWithLocation)
	}

	return suggestionsWithLocation, err
}

func GetSuggestionsByTouristUsername(touristUsername string, db *gorm.DB) (suggestions []models.Suggestion, err error) {
	err = db.Model(&models.Suggestion{}).Where("tourist_username = ?", touristUsername).Find(&suggestions).Error
	if err != nil {
		return nil, err
	}
	return suggestions, nil
}

func GetAllSuggestionsWithLocationByTouristUsername(touristUsername string, db *gorm.DB) (suggestionsWithLocation []models.SuggestionWithLocation, err error) {
	// find all suggestions by tourist username in the database
	var suggestions []models.Suggestion
	suggestions, err = GetSuggestionsByTouristUsername(touristUsername, db)

	if err != nil {
		return nil, err
	}

	// get all suggestions with location
	for _, suggestion := range suggestions {
		suggestionWithLocation, err := GetSuggestionWithLocationBySuggestionId(suggestion.SuggestionId, db)
		if err != nil {
			return nil, err
		}
		suggestionsWithLocation = append(suggestionsWithLocation, *suggestionWithLocation)
	}

	return suggestionsWithLocation, err
}
