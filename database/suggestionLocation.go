package database

import (
	"DebTour/models"

	"gorm.io/gorm"
)

func GetSuggestionLocationBySuggestionId(suggestionId uint, db *gorm.DB) (suggestionLocation *models.SuggestionLocation, err error) {
	err = db.Model(&models.SuggestionLocation{}).Where("suggestion_id = ?", suggestionId).First(&suggestionLocation).Error
	if err != nil {
		return nil, err
	}
	return suggestionLocation, nil
}

func CreateSuggestionLocation(suggestionLocation *models.SuggestionLocation, db *gorm.DB) (err error) {
	db.SavePoint("BeforeCreateSuggestionLocation")
	err = db.Create(suggestionLocation).Error
	if err != nil {
		db.RollbackTo("BeforeCreateSuggestionLocation")
		return err
	}
	return nil
}

func UpdateSuggestionLocation(suggestionId uint, suggestionLocation *models.SuggestionLocation, db *gorm.DB) (err error) {
	db.SavePoint("BeforeUpdateSuggestionLocation")
	err = db.Model(&models.SuggestionLocation{}).Where("suggestion_id = ?", suggestionId).Updates(suggestionLocation).Error
	if err != nil {
		db.RollbackTo("BeforeUpdateSuggestionLocation")
		return err
	}
	return nil
}

func DeleteSuggestionLocation(suggestionId uint, db *gorm.DB) (err error) {
	db.SavePoint("BeforeDeleteSuggestionLocation")
	err = db.Where("suggestion_id = ?", suggestionId).Delete(&models.SuggestionLocation{}).Error
	if err != nil {
		db.RollbackTo("BeforeDeleteSuggestionLocation")
		return err
	}
	return nil
}
